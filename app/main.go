package main

import (
	"fmt"
	"strings"
)



func main() {

	for{

		var userInput string
	
		fmt.Print("$ ");

		
	
		fmt.Scan(&userInput)
     
		 
		parts:=strings.SplitN(userInput," ",2)
      command:=parts[0]
		if command=="exit"{
			 break
		}else if command=="echo"{
			 fmt.Printf("%s\n",parts[1]);
		}else{

			fmt.Printf("%s: command not found\n",userInput)
		}
	}

}
