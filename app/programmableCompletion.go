package main

import "fmt"


func complete(args []string){

	
	   if len(args)<1{
			  fmt.Println("Complete expected an argument ")
			  return
		}
		
		if args[0]=="-p"{
			   if len(args)<2{
					  
					  fmt.Println("flag -p expected a specification")
					  return
				}

				fmt.Printf("complete: %s: no completion specification\n",args[1])
		}

}