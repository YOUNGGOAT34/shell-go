package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)


type JobStatus int

type Job struct{
	    jobNumber int
		 PID int
		 status JobStatus
		 command string
}


const(
     
	   Running JobStatus=iota

)


func (status JobStatus) jobStatusToString() string{
	    switch status{
		 		case Running:
						return "Running"
				default:
					  return "unknown"
		 }
}


var nextJobNumber=1
var jobs []Job

func startBackGroundJob(command string,args []string) bool{
	  cmd:=exec.Command(command,args[:len(args)-1]...)

	  cmd.Stdout=os.Stdout
	  cmd.Stderr=os.Stderr

	  err:=cmd.Start()

	  if err!=nil{
		 return false
	  }

	
	  Job:=Job{
		     jobNumber:nextJobNumber,
			  PID: cmd.Process.Pid,
			  status: Running,
			  command:command+" "+strings.Join(args," "),
	  }


	  jobs = append(jobs, Job)

	  nextJobNumber++

	  fmt.Printf("[%d] %d\r\n",Job.jobNumber,Job.PID)
      
    return true
	    
}

func showJobs(){
	  for _,job:= range jobs{

		  fmt.Printf("[%d]+  %-24s %s\n",job.jobNumber,job.status.jobStatusToString(),job.command)
	  }
}