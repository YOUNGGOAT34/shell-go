package main

import (
	 "strings"
	
	
)

func parseUserInput(userInput []rune,redirect *Redirect) []string{
	   var args []string
      
		var currentArg strings.Builder;
		inSingleQuotes:=false
		inDoubleQuotes:=false


		redirect.stdout=false
		redirect.stderr=false
		redirect.fileName=""
      

	

		for i:=0;i<len(userInput);i++{
             char:=userInput[i]

				 //handle backslash

			    if char=='\\' && !inSingleQuotes && !inDoubleQuotes{
					    if i<len(userInput)-1{

							 currentArg.WriteRune(userInput[i+1])
							 i++
						 }

						 continue
						 
				 }else if char=='\\' && inDoubleQuotes{
					   
					     //at this stage I should only escape " and \ ,,everything else should be treated as literal characters  
						  
						  if i<len(userInput)-1 && (userInput[i+1]=='"' || userInput[i+1]=='\\'){

										currentArg.WriteRune(userInput[i+1])
									   i++
						}else{
							currentArg.WriteRune(char)
						}
						  continue
				 }

				 if char=='>' || (char =='1' && i<len(userInput)-1 && userInput[i+1]=='>'){

					   redirect.stdout=true
						if currentArg.Len()>0{
							  args=append(args,currentArg.String())
							  currentArg.Reset()
						}

						if char=='1' {
							 i++
						}

						if  i<len(userInput)-1 && userInput[i+1]=='>'{
							  i++
						}
						
						continue

				 }else if char=='2' && i<len(userInput)-1 && userInput[i+1]=='>'{
					    redirect.stderr=true
                    
						 if currentArg.Len()>0{
							  args=append(args,currentArg.String())
							  currentArg.Reset()
						}

						i++

						//for appending

						if i<len(userInput) && userInput[i]=='>'{
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

			  if redirect.stdout || redirect.stderr{
					  redirect.fileName=currentArg.String()
				 }else{
                  
					args=append(args,currentArg.String())
					  
				 }

			  
		}

		return args
}