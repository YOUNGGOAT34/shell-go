package main

import (
	"fmt"
)



func main() {

	for{

		var userInput string
	
		fmt.Print("$ ");
	
		fmt.Scan(&userInput)
		if userInput=="exit"{
			 break
		}
		fmt.Printf("%s: command not found\n",userInput)
	}

}
