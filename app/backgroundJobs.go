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
		Done

)


func (status JobStatus) jobStatusToString() string{
	    switch status{
		 		case Running:
						return "Running"
				case Done:
					   return "Done"
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

	

	  //This routine will update the job's status to done ,once the job is done running

	  go func(jobNumber int){
          cmd.Wait()
         
			 for i:=range jobs{
				  if jobs[i].jobNumber==jobNumber{
					   
					  jobs[i].status=Done
				  }
			 }

	  }(Job.jobNumber)

	  nextJobNumber++

	  fmt.Printf("[%d] %d\r\n",Job.jobNumber,Job.PID)
   
    return true
	    
}


func showJobs(){
	  for i:=0;i<len(jobs);{
        indicator:=" "
        if i==len(jobs)-1{
           indicator="+"
		  }else if len(jobs)>=2 && i==len(jobs)-2{
			   indicator="-"
		  }

		command:=jobs[i].command

		if jobs[i].status==Done{
			 /*
			    command is guaranteed to have a trailing & ,because only background jobs are stored in the jobs slice 
				 Therefore it is safe to access len(command)-1
			 */
			  command=command[:len(command)-1]
			  
		}

		fmt.Printf("[%d]%s  %-24s %s\n",jobs[i].jobNumber,indicator,jobs[i].status.jobStatusToString(),command)

		if jobs[i].status==Done{
			  jobs=append(jobs[:i],jobs[i+1:]... )
			  
		}else{
			 i++
		}

	  }
}