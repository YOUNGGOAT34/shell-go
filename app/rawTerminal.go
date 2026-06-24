package main

import (
	"errors"
	"fmt"
	"io"
	"os"
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



func findLastSpace(userInput []rune) int{
	   lastSpaceIndex:=-1
	   for i,r :=range userInput{
			   if r==' ' && i<len(userInput)-1{
					 
					 lastSpaceIndex=i
				}
		}
	   return lastSpaceIndex
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

									 fmt.Printf("^C\r\n")
									
									_break=true
									
								case '\t':
									
									  
									   tab_count++

										/*
										    If the input is something like : cat file/file1

											 I don't want to pass the entire thing to the autocompletion ,I only want to pass the part after the last space
										*/

										lastSpace:=findLastSpace(userInput)

										var inputToAutocomplete []rune
										var searchInDir bool

										if lastSpace != -1{

                                 inputToAutocomplete=userInput[lastSpace+1:]
											searchInDir=true

										}else{
											 inputToAutocomplete=userInput
											 if userInput[len(userInput)-1]==' '{
												    
												    searchInDir=true
											 }else{
												  searchInDir=false
											 }
											 
										}

										/*
										  I wanna pass the full user input to the autocomplete function so that for programmable completion
										  It will be used to form ,command,current word and previous word
										*/

                              matches:=autocomplete(inputToAutocomplete,userInput,searchInDir)
                              
										switch tab_count {
											
                                 case 1:
													
													if len(matches)<1 || matches==nil || len(matches)>1{
															
															
															if len(matches)>1{

                                                 	/*
																	If there were multiple matches and they share a common prefix ,autocomplete with the longest common prefix,
																	otherwise ring a bell
																	the longestCommonPrefix function returns the longest common prefix in the matches
																	
																	*/

																lcp:=longestCommonPrefix(matches)

																if lcp>0{

																	spaceIndex:=findLastSpace(userInput)

																	if spaceIndex==-1{

																				if lcp>len(userInput){
																					userInput=matches[0][:lcp]
																					tab_count=0
																				}else{
																					fmt.Print("\a") 
																				}
																	}else if spaceIndex>0{
																		  
																				/*
																					At this point it might be a directory or a just a file within a directory
																					we can tell this by finding if the user input  had a /
																					the findLastSlash function finds the position of the last / in a given path 
																					if there was no / it returns 0 ,this will indicate that it was either a file in the parent directory
																					or a directory in the parent directory
																				*/
						
																				lastSlashIndex:=findLastSlash(userInput)

																				/* 
																					   If this prefix is longer than the current input ,then autcomplete ,otherwise ring a bell
																					*/
																				
						                                          prefix:=matches[0][:lcp]

																				if lastSlashIndex==0{

																				

																					if len(prefix)>len(userInput[spaceIndex+1:]){

                                                                      userInput=append(userInput[:spaceIndex+1],prefix... )
																							 tab_count=0
                                                                     
																					}else{
																						  fmt.Print("\a") 
																					}
																					   
																						
																				}else{

																					if len(prefix)>len(userInput[lastSlashIndex+1:]){
																						userInput=append(userInput[:lastSlashIndex+1],prefix...)
																						tab_count=0
																					}else{
																						  fmt.Print("\a") 
																					}

																				}
																	}
																	
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
																the findLastSpace function returns -1 if there was no space
																the function finds the last space

                                             */

														    spaceIndex:=findLastSpace(userInput)
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

																			if userInput[len(userInput)-1]==' '{
																				userInput=append(userInput, matches[0]...)
																			}else{
																				
																				userInput=matches[0]
																			}
																		}
																		
																		tab_count=0
																}

													case 2:
															printMatches(matches)
															tab_count=0
											}

								          case 127:

												userInput=userInput[:len(userInput)-1]

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