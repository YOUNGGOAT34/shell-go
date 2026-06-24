package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
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




func autcompletePrammableCompletion(userInput []rune) [][]rune{
   
	   var matches [][]rune

		command,currentWord,previousWord:=completionArgs(userInput)

		if path,exists:=completions[strings.TrimSpace(string(command))];exists{
			        
			        cmd:=exec.Command(path,command,currentWord,previousWord)

					  cmd.Env=append(cmd.Env, 
					    "COMP_LINE="+string(userInput),
						 "COMP_POINT="+strconv.Itoa(len(string(userInput))),
					)


					output,err:=cmd.Output()

				


					  
					
					  if err==nil{

						        if len(output)!=0{

									 lines:=strings.Split(string(output),"\n")

									 for _,line:=range lines{
										 
										 match:=[]rune(strings.TrimSpace(string(line)))
	  
										 match=append(match, ' ')
	  
										 matches=append(matches, match)
									 }

								  }
                    
									 
					  }

					  
		}
		
		return matches
}




/*
   SearchInDir is a flag to tell the program where it should search for matches
	if true ,it means search in the current directory or subdirectories else search in path and builtins
*/
func autocomplete(currentInput []rune,fullInput []rune,searchInDir bool) ([][]rune){


    
	var matches[][] rune
    
	 //autcomplete a programmable completion
	
	 matches=autcompletePrammableCompletion(fullInput)

	 if len(matches)>0{
		 return matches
	 }

	 if searchInDir{

		     lastSlash:=findLastSlash(currentInput)
			  
		     if lastSlash>0{
                matches=searchInDirectory(currentInput[lastSlash+1:],string(currentInput[:lastSlash]))
			  }else{
				  /*
				      If the user typed something like : du<space><tab>
						This is the current input ,therefore we whave an empty prefix
						This is what we want to pass to the search directory so that we return the first file/directory as our match
				  */
				  if currentInput[len(currentInput)-1]==' '{
					  
					    matches=searchInDirectory([]rune(""),".")
				  }else{

					  matches=searchInDirectory(currentInput,".")
				  }
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

