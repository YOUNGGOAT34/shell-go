package main

import (
	
	"os"
	"fmt"
)


func isExecutable(path string) bool{
	  info,err:=os.Stat(path)

	  if err!=nil{
		 return false
	  }

	  if info.IsDir(){
		 return false
	  }

	  return info.Mode() & 0111 !=0


}


func searchInCurrentDirectory(userInput []rune) [][]rune{

	   
	    entries,err:=os.ReadDir(".")

		 var matches [][]rune
       
		 if err!=nil{
			    fmt.Fprintln(os.Stderr,"Error: ",err)
				 return matches
		 }

		 for _,entry :=range entries{
			    if entry.IsDir(){
					  continue
				 }

				 fileName:=[]rune(entry.Name())

				 if hasPrefixRune(fileName,userInput){
					  matches=append(matches, []rune(fileName))
					  
				 }
		 }

		 return matches

}

