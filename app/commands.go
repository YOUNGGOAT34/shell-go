package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)


type Redirect struct{
	  stdout bool
	  stderr bool
	  fileName string
	 
}






func createRedirectFile(fileName string) *os.File{
	  file,err:=os.OpenFile(fileName,os.O_APPEND | os.O_CREATE | os.O_WRONLY,0644)

	  if err!=nil{
		   fmt.Printf("Error opening/creating redirect file :%v",err)

			return nil
			
	  }

	  return file
}


func isInbuilt(command string) bool{
	    inbuilts:=map[string] bool{
			"exit":true,
			"type":true,
			"echo":true,
			"pwd":true,
			"complete":true,
			"jobs":true,
		}
		return inbuilts[command]
}

var redirect Redirect

func execute(userInput []rune) bool{

	   if len(userInput)<1{
			 return true
		}


	  args:=parseUserInput(userInput,&redirect)

      command:=args[0]

		switch command {
				case "exit":
					return false
				case "echo":
					
					handleEcho(args[1:])
				case "type":
					      if len(args)<2{
                            fmt.Printf("type expected an argument\n")
									 return true
							}
							if isInbuilt(args[1]){
								fmt.Printf("%s is a shell builtin\n",args[1]);
							}else{
                          
									handleType(args[1])		
							}
				case "pwd":
					   
					   printWorkingDirectory()

				case "cd":
					    if len(args)<2{
                      fmt.Println("cd expected an argument")
						 }else{

							 changeDirectory(args[1])
						 }

				case "complete":
					  
					  complete(args[1:])

				case "jobs":

				default:
					if !runProgram(command,args[1:]){
              
						fmt.Printf("%s: command not found\n",command)
					}
				}


		return true

}


func handleEcho(args []string){
       oldStdout:=os.Stdout
		 oldStderr:=os.Stderr
	    if redirect.stdout{
			  file:=createRedirectFile(redirect.fileName)
			  defer file.Close()
			  
			  os.Stdout=file
			 
		 }

		  if redirect.stderr{
			  file:=createRedirectFile(redirect.fileName)
			  defer file.Close()
			  
			  os.Stderr=file
			 
		 }
	 
	    
	    if len(args)>0{
			 fmt.Println(strings.Join(args," "))
		 }else{
			 fmt.Println()
		 }

		  os.Stdout=oldStdout
		  os.Stderr=oldStderr
}


func handleType(cmd string){
	  
	    pathEnv:=os.Getenv("PATH");
		 dirs:=filepath.SplitList(pathEnv)
		 found:=false
      
		for _,dir :=range dirs{
			fullPath:=filepath.Join(dir,cmd)

			if isExecutable(fullPath){
					fmt.Printf("%s is %s\n",cmd,fullPath)
					found=true
					return
					
			}

		}
		
		if !found{

			fmt.Printf("%s: not found\n",cmd)
		}

}

func runProgram(command string,args []string) bool{

      if len(args)>=1 && args[len(args)-1]=="&"{
			 return startBackGroundJob(command,args[:len(args)-1])
		}
	   
		cmd:=exec.Command(command,args...)

		oldStdout:=os.Stdout
		oldStderr:=os.Stderr

		var file *os.File
		if redirect.stdout || redirect.stderr {
         file=createRedirectFile(redirect.fileName)
			

			if file==nil{
				  
				 cmd.Stdout=os.Stdout
				 cmd.Stderr=os.Stderr
			}else{

				 defer file.Close()

				   if redirect.stdout {
      
           			cmd.Stdout=file

						}else{
							cmd.Stdout=os.Stdout
							
						}
					if redirect.stderr {
        

							cmd.Stderr=file
							
					
						}else{
							cmd.Stderr=os.Stderr
							
						}

			}
    
		}else{
			cmd.Stdout=os.Stdout
			cmd.Stderr=os.Stderr
		}


		cmd.Stdin=os.Stdin
	
		
		err:=cmd.Run()



		if err !=nil{
			    if errors.Is(err,exec.ErrNotFound){
					  return false
				 }
		}

		os.Stdout=oldStdout
		os.Stderr=oldStderr


		return true
}