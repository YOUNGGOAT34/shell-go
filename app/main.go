package main

import (
	"fmt"
)



func main() {

	for{

		var userInput string
	
		fmt.Print("$ ");
	
		fmt.Scan(&userInput)
		fmt.Printf("%s: command not found\n",userInput)
	}

}
