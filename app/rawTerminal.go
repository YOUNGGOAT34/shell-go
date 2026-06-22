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



func hasPrefixRune(fullCommand []rune,currentInput []rune) bool{
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



func executableAutocompletion(userInput *[]rune) [][]rune{
	   pathEnv:=os.Getenv("PATH")

		var matches [][]rune

		currentInput:=*userInput

		dirs:=filepath.SplitList(pathEnv)

		for _,dir:=range dirs{
			   entries,err:=os.ReadDir(dir)

				if err!=nil{
					 continue
				}

				for _,entry :=range entries{
					   if !entry.IsDir(){

							    if hasPrefixRune([]rune(entry.Name()),currentInput){
									    matches=append(matches,[]rune(entry.Name()))
										 
								 }
							   
						}
				}
		}

		return matches
}


func autocomplete(userInput *[]rune) ([][]rune){

    currentInput:=*userInput
	 var matches[][] rune

   
	 if len(currentInput)==0{
		 return nil
	 }

	 for _,builtin :=range builtins{
		   if hasPrefixRune(builtin,currentInput){
				  match:=make([]rune,len(builtin))
				  copy(match,builtin)
              matches=append(matches,match)
				  continue
			}
	 }

	 if len(matches)>0{
		  return matches
	 }

   matches=append(matches,executableAutocompletion(userInput)...)

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
                              matches:=autocomplete(&userInput)
										
										if tab_count==1{
											 if len(matches)<1 || matches==nil || len(matches)>1{
												   fmt.Print("\a")
											 }else if len(matches)==1{
												    userInput=matches[0]
													 userInput=append(userInput, ' ')
													 tab_count=0
											 }
										}else if tab_count==2{
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