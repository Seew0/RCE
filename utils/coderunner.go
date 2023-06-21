package utils

import (
	"log"
	"os/exec"
	"rce/models"
)

func CodeRunner(inputChan chan models.Request, outputChan chan models.Response) {
	for job := range inputChan {
		newExecution := ExecutionAdapter(job.ReqID,job.Language,job.Code,job.CodeInput)

		file, err := newExecution.TempFile()
		if err != nil {
			// something error
		}

		var cmd *exec.Cmd

		if job.Language == "c" || job.Language == "cpp" || job.Language == "java" {
			var buildresp *models.Response
			cmd, buildresp, err = newExecution.ExecuteCompilable(file)
			if len(buildresp.Errors) != 0 {
				outputChan <- *buildresp
				continue
			}
			if err != nil {
				log.Println(err)
			}
		}

		if job.Language == "python3" || job.Language == "go" || job.Language == "bash" {
			cmd, err = newExecution.ExecuteInterpreted(file)
			if err != nil {
				log.Println(err)
			}
		}

		resp, err := newExecution.Run(cmd, file)
		if err != nil {
			log.Println(err)
		}

		outputChan <- *resp
	}
}
