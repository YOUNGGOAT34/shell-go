package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



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
		}else{

			fmt.Printf("%s: command not found\n",userInput)
		}
	}

}
