package main

import (
	"fmt"
	"strings"
)


var completions =make(map[string] string)


func completionArgs(userInput []rune)(command,currentWord,previousWord string){

	 words:=strings.Fields(string(userInput))

	 endsWithSpace:=len(userInput)>0 && userInput[len(userInput)-1]==' '

	 if len(words)==0{
		  return "","",""
	 }

	 command=words[0]



	 if len(words)==1{
		  return command,"",""
	 }


	 if endsWithSpace{
		  currentWord=""

		  if len(words)>=2{
			  previousWord=words[len(words)-1]
		  }
	 }else{
		  
		   currentWord=words[len(words)-1]

			if len(words)>=3{
				 previousWord=words[len(words)-2]
			}
	 }



	 return command,currentWord,previousWord


}


func complete(args []string){

	
	   if len(args)<1{
			  fmt.Println("Complete expected a flag i.e -p ")
			  return
		}
		
		if args[0]=="-p"{
			   if len(args)<2{
					  
					  fmt.Println("flag -p expected a specification")
					  return
				}

				printCompletion(args[1])


		}

		if args[0]=="-C"{
			   if len(args)<3{
					 fmt.Println("flag -C expects two arguments:path and specification name")
					 return
				}

				registerCompletion(args[1],args[2])
		}

}


func printCompletion(completionName string){
      path,exists:=completions[completionName]

		if exists{
			  fmt.Printf("complete -C '%s' %s\n",path,completionName)
		}else{
           fmt.Printf("complete: %s: no completion specification\n",completionName)
		}
}


func registerCompletion(path string,completionName string){
	    completions[completionName]=path
}