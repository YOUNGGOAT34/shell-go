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
var freeJobNumbers []int
var jobs []Job

func startBackGroundJob(command string,args []string) bool{
	  cmd:=exec.Command(command,args[:len(args)-1]...)

	  cmd.Stdout=os.Stdout
	  cmd.Stderr=os.Stderr

	  err:=cmd.Start()

	  if err!=nil{
		 return false
	  }



	  var jobNumber int

	  if len(jobs)>0{

		   if len(freeJobNumbers)>0{
				 jobNumber=freeJobNumbers[0]
				 freeJobNumbers=append(freeJobNumbers,freeJobNumbers[1:]... )
			}else{
				 jobNumber=nextJobNumber
				 nextJobNumber++
			}

	  }else{
		  jobNumber=nextJobNumber
		  nextJobNumber++
	  }

	
	  Job:=Job{
		     jobNumber:jobNumber,
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

	  

	  fmt.Printf("[%d] %d\r\n",Job.jobNumber,Job.PID)
   
    return true
	    
}



/*
     reapBeforePrompt determines how completed jobs are handled.

   false (jobs command):
     - Display all jobs.
     - Reap completed jobs after displaying them.

   true (before prompt):
     - Display only completed jobs.
     - Reap them immediately.
*/

func showJobs(reapBeforePrompt bool){
	
	  for i:=0;i<len(jobs);{
        indicator:=" "
        if i==len(jobs)-1{
           indicator="+"
		  }else if len(jobs)>=2 && i==len(jobs)-2{
			   indicator="-"
		  }

		command:=jobs[i].command

		if jobs[i].status==Running && reapBeforePrompt{
			  i++
			  continue
		}

		if jobs[i].status==Done{
			 /*
				command is guaranteed to end with '&' because only
				background jobs are stored in the jobs slice.
				Therefore command[:len(command)-1] is safe.
			*/
			  command=command[:len(command)-1]
			  
		}

		fmt.Printf("[%d]%s  %-24s %s\r\n",jobs[i].jobNumber,indicator,jobs[i].status.jobStatusToString(),command)

		if jobs[i].status==Done{
			 
           
			  freeJobNumbers=append(freeJobNumbers,jobs[i].jobNumber)
			   jobs=append(jobs[:i],jobs[i+1:]... )

			  if len(jobs)==0{
				  nextJobNumber=1
			  }
			  
		}else{
			 i++
		}

	  }
}

