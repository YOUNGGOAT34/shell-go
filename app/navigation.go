package main

import (
	 "fmt"
	 "os"
)


func printWorkingDirectory(){
      dir,err:=os.Getwd()

		if err!=nil{
			  fmt.Println("Failed to get the current working directory");
			  return
		}

		fmt.Println(dir)

}


func changeDirectory(path string){
	    os.Chdir(path)
}