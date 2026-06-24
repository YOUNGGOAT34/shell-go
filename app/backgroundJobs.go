package main

import (
	"fmt"
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