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





func searchInDirectory(userInput []rune,path string) [][]rune{
       
		 entries,err:=os.ReadDir(path)
		
		 var matches [][]rune
       
		 if err!=nil{
			   
			    fmt.Fprintln(os.Stderr,"Error: ",err)
				 return matches
		 }

		 for _,entry :=range entries{
			    

				 fileOrDirectoryName:=[]rune(entry.Name())

				 

				 if hasPrefixRune(fileOrDirectoryName,userInput){
                   /*
                      when the user presses tab ,autocompletion for a file name should ,the match should be prefixed with a space ,
							 But for a directory it should be prefixed with a /

						 */
					 
                 if entry.IsDir(){
                    fileOrDirectoryName=append(fileOrDirectoryName, '/')
					  }else{
                         fileOrDirectoryName=append(fileOrDirectoryName, ' ')
						  }	

						  matches=append(matches,fileOrDirectoryName)
					  
				 }
		 }

		 return matches

}

