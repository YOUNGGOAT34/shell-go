package main

import "fmt"


var completions =make(map[string] string)


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