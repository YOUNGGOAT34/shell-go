package main

import (
	"fmt"
	"os"
	"io"
    "errors"
	 "unicode/utf8"

	"golang.org/x/term"
)

var builtins=[][]rune{
      []rune("exit"),
		[]rune("echo"),
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


func autocomplete(userInput *[]rune) bool{

    currentInput:=*userInput
	 
	 for _,builtin :=range builtins{
		   if hasPrefixRune(builtin,currentInput){
				  *userInput=builtin
				  *userInput=append(*userInput,' ')
				  return true
			}
	 }

	 return false
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
										if !autocomplete(&userInput){
											fmt.Printf("\a")
										}
								default:
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