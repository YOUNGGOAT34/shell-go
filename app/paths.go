package main

import (
	 "os"
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