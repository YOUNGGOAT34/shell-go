package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)




func printMatches(matches [][]rune){
	 
	  if len(matches)<1{
		 return
	  }

	  var matchStrings []string

	  for _,match :=range matches{
		     
		   str:=strings.TrimSpace(string(match))

			matchStrings=append(matchStrings, str)
	  }


	  sort.Strings(matchStrings)

	  fmt.Printf("\r\n%s\r\n",strings.Join(matchStrings,"  "))
	  os.Stdout.Sync()
}





/*
   SearchInDir is a flag to tell the program where it should search for matches
	if true ,it means search in the current directory or subdirectories else search in path and builtins
*/
func autocomplete(currentInput []rune,searchInDir bool) ([][]rune){


   

	 var matches[][] rune


	 if searchInDir{

		     lastSlash:=findLastSlash(currentInput)
			  
		     if lastSlash>0{
                matches=searchInDirectory(currentInput[lastSlash+1:],string(currentInput[:lastSlash]))
			  }else{
				 
				  matches=searchInDirectory(currentInput,".")
			  }
			  return matches
			  
	}

   
	 if len(currentInput)==0{
		 return nil
	 }

	 for _,builtin :=range builtins{
		   if hasPrefixRune(builtin,currentInput){
				  match:=make([]rune,len(builtin))
				  copy(match,builtin)
				  match=append(match, ' ')
              matches=append(matches,match)
				  continue
			}
	 }

	 if len(matches)>0{
		  return matches
	 }

	


   matches=append(matches,executableAutocompletion(currentInput)...)

	 return matches
}




func executableAutocompletion(currentInput []rune) [][]rune{
	   pathEnv:=os.Getenv("PATH")

		var matches [][]rune

		

		dirs:=filepath.SplitList(pathEnv)

		for _,dir:=range dirs{
			   entries,err:=os.ReadDir(dir)

				if err!=nil{
					 continue
				}

				for _,entry :=range entries{
					   if !entry.IsDir(){

							    if hasPrefixRune([]rune(entry.Name()),currentInput){
									    match:=[]rune(entry.Name())
										 match=append(match, ' ')
									    matches=append(matches,match)
										 
								 }
							   
						}
				}
		}

		return matches
}



func longestCommonPrefix(matches [][]rune) int{
   base:=matches[0]
	

	 longestCommonPrefix:=0
	 
    for index,char:= range base{
		    for _,match:=range matches[1:]{
				   if index>=len(match) || (index<len(match) && match[index] !=char){
						  return longestCommonPrefix
					}
			 }

			 longestCommonPrefix+=1
	 }

	 return longestCommonPrefix
}


func hasPrefixRune(fullCommand []rune,currentInput []rune) bool{
   

	  if len(currentInput)==0{
		 return true
	  }

	  if len(currentInput)>len(fullCommand){
		 return false
	  }

	  for i:=range currentInput{
		    if fullCommand[i]!=currentInput[i]{
				 return false
			 }
	  }


	  return true
}

