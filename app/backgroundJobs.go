package main

import (
	"fmt"
	"os"
	"os/exec"
)



type Job struct{
	    jobNumber int
		 PID int
}





var nextJobNumber=1
var jobs []Job

func startBackGroundJob(command string,args []string) bool{
	  cmd:=exec.Command(command,args...)

	  cmd.Stdout=os.Stdout
	  cmd.Stderr=os.Stderr

	  err:=cmd.Start()


	  if err!=nil{
		 return false
	  }

	  

	  Job:=Job{
		     jobNumber:nextJobNumber,
			  PID: cmd.Process.Pid,
	  }


	  jobs = append(jobs, Job)

	  nextJobNumber++

	  fmt.Printf("[%d] %d\r\n",Job.jobNumber,Job.PID)
      
    return true
	    
}