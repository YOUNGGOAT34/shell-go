package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"sort"
	"os"
	"path/filepath"
	"unicode/utf8"

	"golang.org/x/term"
)

var builtins=[][]rune{
      []rune("exit"),
		[]rune("echo"),
		[]rune("pwd"),
		[]rune("type"),
		
}




func findLastSlash(userInput []rune) int{
	 
	   lastIndex:=0

		for index,char :=range userInput{
			   if char=='/'{
					  lastIndex=index
				}
		}
      
		return lastIndex

}



func findFirstSpace(userInput []rune) int{
	   for i,r :=range userInput{
			   if r==' '{
					 return i
				}
		}
	   return -1
}



func longestCommonPrefix(matches [][]rune) int{
   base:=matches[0]
	

	 longestCommonPrefix:=0
	 
    for index,char:= range base{
		    for _,match:=range matches[1:]{
				   if index>=len(match) || (index<len(match) && match[index] !=char){
						  return longestCommonPrefix
					}
			 }

			 longestCommonPrefix+=1
	 }

	 return longestCommonPrefix
}


func hasPrefixRune(fullCommand []rune,currentInput []rune) bool{
   

	  if len(currentInput)==0{
		 return true
	  }

	  if len(currentInput)>len(fullCommand){
		 return false
	  }

	  for i:=range currentInput{
		    if fullCommand[i]!=currentInput[i]{
				 return false
			 }
	  }


	  return true
}



func executableAutocompletion(currentInput []rune) [][]rune{
	   pathEnv:=os.Getenv("PATH")

		var matches [][]rune

		

		dirs:=filepath.SplitList(pathEnv)

		for _,dir:=range dirs{
			   entries,err:=os.ReadDir(dir)

				if err!=nil{
					 continue
				}

				for _,entry :=range entries{
					   if !entry.IsDir(){

							    if hasPrefixRune([]rune(entry.Name()),currentInput){
									    match:=[]rune(entry.Name())
										 match=append(match, ' ')
									    matches=append(matches,match)
										 
								 }
							   
						}
				}
		}

		return matches
}


func autocomplete(currentInput []rune) ([][]rune){


   

	 var matches[][] rune

  
	 spaceIndex:=findFirstSpace(currentInput)


	 if spaceIndex!=-1{
		     lastSlash:=findLastSlash(currentInput)
			  
		     if lastSlash>0{
                matches=searchInDirectory(currentInput[lastSlash+1:],string(currentInput[spaceIndex+1:lastSlash]))
			  }else{
				  matches=searchInDirectory(currentInput[spaceIndex+1:],".")
			  }
			  return matches
			  
	}

   
	 if len(currentInput)==0{
		 return nil
	 }

	 for _,builtin :=range builtins{
		   if hasPrefixRune(builtin,currentInput){
				  match:=make([]rune,len(builtin))
				  copy(match,builtin)
				  match=append(match, ' ')
              matches=append(matches,match)
				  continue
			}
	 }

	 if len(matches)>0{
		  return matches
	 }

   matches=append(matches,executableAutocompletion(currentInput)...)

	 return matches
}


func printMatches(matches [][]rune){
	  if len(matches)<1{
		 return
	  }

	  var matchStrings []string

	  for _,match :=range matches{
		     
		   str:=strings.TrimSpace(string(match))

			matchStrings=append(matchStrings, str)
	  }


	  sort.Strings(matchStrings)

	  fmt.Printf("\r\n%s\r\n",strings.Join(matchStrings,"  "))
}



func processRawInput() []rune{

	   fd:=int(os.Stdin.Fd())
      
		oldTerminalState,err:=term.MakeRaw(fd)

		if err!=nil{
			 panic(err)
		}

		defer term.Restore(fd,oldTerminalState)

      
    
		var userInput []rune

		buffer:=make([]byte,3)

		fmt.Print("$ ");

		tab_count:=0
       
		for{

			   

			  _break:=false

			  bytesRead,err:=os.Stdin.Read(buffer)



			  if err!=nil{
				   if errors.Is(err,io.EOF){
						  break
					}
				   panic(err)
			  }

			   if bytesRead>0{

					i:=0

					for i<bytesRead{

						char,size:=utf8.DecodeRune(buffer[i:bytesRead])
	               i+=size
						switch char {
								case '\r', '\n':
									
									_break=true
									
								case 3:
									
									_break=true
									
								case '\t':
									  

									   tab_count++

                              matches:=autocomplete(userInput)

										switch tab_count {
											
                                 case 1:
													
													if len(matches)<1 || matches==nil || len(matches)>1{
															
															
															if len(matches)>1{
																lcp:=longestCommonPrefix(matches)
																
																if lcp>len(userInput){
																	userInput=matches[0][:lcp]
																	tab_count=0
																}else{
																	 fmt.Print("\a") 
																}
															}else{
                                                 
																fmt.Print("\a")
															}
													}else if len(matches)==1{


														  /*
														      if there was a space we just want to overwrite the second part

																for example the user types something like : cat mai  
																and the presses tab ,then only part that should be completed is the mai part ,
																not the entire input ,therefore we have to find if there was space inside the user input
																the findFirstSpace function returns -1 if there was no space

                                             */

														    spaceIndex:=findFirstSpace(userInput)
															 

														   if spaceIndex!=-1{

																 /*
                                                      what if the user typed in something like cat path/to/f

																		to autocomplete this ,we can't overwrite the entire second part ,therefore we need to find the last
																		occurrence of /

																		then overwrite everything that follows with our match
																    
																 */

																 

																

																 lastSlash:=findLastSlash(userInput)
                                                  
																 if lastSlash>0{
																	 
																	 userInput=append(userInput[:lastSlash+1],matches[0]...)

																 }else{
																	   userInput=append(userInput[:spaceIndex+1],matches[0]...)
																 }
                                                 
																
                                                
															}else{
																userInput=matches[0]
																
																
															}
															
															tab_count=0
													}
											case 2:
													printMatches(matches)
													tab_count=0
											}
             
                              

								default:
									tab_count=0
									userInput=append(userInput,char)
						}
					}


					fmt.Printf("\r\033[K$ %s",string(userInput))
				   if _break{
						 fmt.Print("\r\n")
					    break
				   }

				 
   
				}
 
		}


		return userInput

}