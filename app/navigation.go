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
	    if path =="~"{
			 home:=os.Getenv("HOME")

			 if home==""{
				 fmt.Println("Home is not set")
				 return
			 }
          
			 os.Chdir(home)

			 return
		 }
	    err:=os.Chdir(path)

		 if err!=nil{
			   fmt.Printf("cd: %s: No such file or directory\n",path)
		 }
}