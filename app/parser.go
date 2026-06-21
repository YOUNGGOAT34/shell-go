package main

import (
	 "strings"
	
)

func parseUserInput(userInput string) []string{
	   var args []string
      
		var currentArg strings.Builder;
		inSingleQuotes:=false
		inDoubleQuotes:=false



		runes:=[]rune(userInput)

		for i:=0;i<len(runes);i++{
             char:=runes[i]

			    if char=='\\' && !inSingleQuotes && !inDoubleQuotes{
					    if i<len(runes){

							 currentArg.WriteRune(runes[i+1])
							 i++
						 }


						 continue
				 }

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
					    if inSingleQuotes {
							// inside single quotes it should just behave like any other character ,no special meaning
							   currentArg.WriteRune(char)
						 }else{
							 inDoubleQuotes=!inDoubleQuotes
						 }
						 continue
				 }

				 currentArg.WriteRune(char)
		}

		if len(currentArg.String())>0{
			  args=append(args,currentArg.String())
		}
		return args
}