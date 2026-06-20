package main
import (
	 "fmt"
	 "strings"
	 "os"
	 "path/filepath"
)



func isInbuilt(command string) bool{
	    inbuilts:=map[string] bool{
			"exit":true,"type":true,"echo":true,
		}
	   

		return inbuilts[command]
}




func execute(userInput string) bool{
	  
	  parts:=strings.SplitN(userInput," ",2)
		
      command:=parts[0]


		switch command {
				case "exit":
					return false
				case "echo":
					handleEcho(parts)
				case "type":
							if isInbuilt(parts[1]){
								fmt.Printf("%s is a shell builtin\n",parts[1]);
							}else{

									handleType(parts)
									
							}
				default:

					fmt.Printf("%s: command not found\n",userInput)
				}


		return true


}

func handleEcho(parts []string){
	    if len(parts)>1{
			 fmt.Println(parts[1])
		 }else{
			 fmt.Println()
		 }
}


func handleType(parts []string){
	  
	    pathEnv:=os.Getenv("PATH");
					 dirs:=filepath.SplitList(pathEnv)
					 found:=false
					 for _,dir :=range dirs{
						   fullPath:=filepath.Join(dir,parts[1])

							if isExecutable(fullPath){
								  fmt.Printf("%s is %s\n",parts[1],fullPath)
								  found=true
								  return
								  
							}

					 }

					 if !found{

						 fmt.Printf("%s: not found\n",parts[1])
					 }

}