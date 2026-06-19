package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"path/filepath"
)


func isExecutable(path string) bool{
	  info,err:=os.Stat(path)

	  if err!=nil{
		 return false
	  }

	  if info.IsDir(){
		 return false
	  }

	  return info.Mode() & 0111 !=0


}



func isInbuilt(command string) bool{
	    inbuilts:=[]string{"exit","type","echo"}
	   for _,cmd :=range inbuilts{
            if cmd==command{
					return true
				}
		}

		return false
}


func main() {

	for{

		
	
		fmt.Print("$ ");

		
	
		reader:=bufio.NewReader(os.Stdin)
     
		userInput,_:=reader.ReadString('\n')

		userInput=strings.TrimSpace(userInput)
	   

		parts:=strings.SplitN(userInput," ",2)
		
      command:=parts[0]

		

		




		if command=="exit"{
			 break
		}else if command=="echo"{
			 fmt.Printf("%s\n",parts[1]);
		}else if command=="type"{
           if isInbuilt(parts[1]){
				  fmt.Printf("%s is a shell builtin\n",parts[1]);
			  }else{
				    pathEnv:=os.Getenv("PATH");
					 dirs:=filepath.SplitList(pathEnv)
					 for _,dir :=range dirs{
						   fullPath:=filepath.Join(dir,parts[1])

							if isExecutable(fullPath){
								  fmt.Printf("%s is %s\n",parts[0],fullPath)
								  return
							}

					 }
				   fmt.Printf("%s: not found\n",parts[1])
			  }
		}else{

			fmt.Printf("%s: command not found\n",userInput)
		}
	}

}
