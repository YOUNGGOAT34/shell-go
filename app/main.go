package main

import (
	"fmt"
)



func main() {

	var userInput string

	fmt.print("$");

	fmt.Scan(&userInput)

	
	fmt.Printf("%s: command not found",userInput)
}
