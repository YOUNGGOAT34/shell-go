package main

import (
	 "strings"
	
)

func parseUserInput(userInput string) []string{
	   var args []string
      
		var currentArg strings.Builder;
		inSingleQuotes:=false
		inDoubleQuotes:=false

		for _,char :=range userInput{
			    if char==' ' &&  !inSingleQuotes && !inDoubleQuotes {
					 
					 if len(currentArg.String())>0{

						 args=append(args,currentArg.String())
						 currentArg.Reset()
					 }

					 continue
				 }

				 if char=='\'' && !inDoubleQuotes{
					  
					  inSingleQuotes=!inSingleQuotes
					  continue
				 }

				 if char=='"'{
                   inDoubleQuotes=!inDoubleQuotes
						 continue
				 }

				 currentArg.WriteRune(char)
		}

		if len(currentArg.String())>0{
			  args=append(args,currentArg.String())
		}
		return args
}