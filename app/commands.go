package main
import (
	 "fmt"
	 "strings"
	 "os"
	 "path/filepath"
	 "os/exec"
)



func isInbuilt(command string) bool{
	    inbuilts:=map[string] bool{
			"exit":true,"type":true,"echo":true,"pwd":true,
		}
		return inbuilts[command]
}

func execute(userInput string) bool{

	   if len(userInput)<1{
			 return true
		}
	  
	  parts:=strings.SplitN(userInput," ",2)
		
      command:=parts[0]
      
		var arguments string

		if len(parts)>1{
           arguments=parts[1]
		}

      
		args:=parseUserInput(arguments)

      
		switch command {
				case "exit":
					return false
				case "echo":

					handleEcho(args)
				case "type":
					      if len(args)<1{
                            fmt.Printf("type expected an argument\n")
									 return true
							}
							if isInbuilt(args[0]){
								fmt.Printf("%s is a shell builtin\n",args[0]);
							}else{

									handleType(args[0])		
							}
				case "pwd":
					   
					   printWorkingDirectory()

				case "cd":
					    if len(args)<1{
                      fmt.Println("cd expected an argument")
						 }else{

							 changeDirectory(parts[1])
						 }

				default:
					if !runProgram(command,args){

						fmt.Printf("%s: command not found\n",command)
					}
				}


		return true

}


func handleEcho(args []string){
	    
	    if len(args)>0{
			 fmt.Println(strings.Join(args," "))
		 }else{
			 fmt.Println()
		 }
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
	  
      

		cmd:=exec.Command(command,args...)

		cmd.Stderr=os.Stderr
		cmd.Stdin=os.Stdin
		cmd.Stdout=os.Stdout

       
		err:=cmd.Run()

		return err==nil
}