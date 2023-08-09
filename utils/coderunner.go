package utils

import (
	"log"
	"os/exec"
	"rce/models"
)

func CodeRunner(req models.Request, outputChan chan models.Response) {

	newExecution := ExecutionAdapter(req.ReqID, req.Language, req.Code, req.CodeInput)

	file, err := newExecution.TempFile()
	if err != nil {
		log.Fatalf("PHATA HAI \n%v", err)
	}

	var cmd *exec.Cmd

	if req.Language == "c" || req.Language == "cpp" || req.Language == "java" {
		var buildresp *models.Response
		cmd, buildresp, err = newExecution.ExecuteCompilable(file)
		if len(buildresp.Errors) != 0 {
			outputChan <- *buildresp
			return
		}
		if err != nil {
			log.Println("idhar hai phatta hua ", err)
		}
	}
	// fmt.Println(req)

	if req.Language == "python3" || req.Language == "go" || req.Language == "bash" {
		cmd, err = newExecution.ExecuteInterpreted(file)
		if err != nil {
			log.Println("execute ni hua ", err)
		}
	}

	resp, err := newExecution.Run(cmd, file)
	if err != nil {
		log.Printf("ni hua %v   resp %v", err, resp)
	}

	outputChan <- *resp
}
