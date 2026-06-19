package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



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
				   fmt.Printf("%s: not found\n",userInput)
			  }
		}else{

			fmt.Printf("%s: command not found\n",userInput)
		}
	}

}
