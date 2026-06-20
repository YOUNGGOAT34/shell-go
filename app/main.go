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
	   
		if !execute(userInput){
			 break
		}
     
		
	}

}
