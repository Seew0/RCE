package utils

import (
	"bytes"
	"errors"
	"os"
	"os/exec"

	"rce/models"
)

type Execution struct {
	CodeID  string
	Language string
	Code    string
	Input   string
}

func ExecutionAdapter(codeid string, Language string, code string, input string) *Execution {
	return &Execution{CodeID: codeid, Language: Language, Code: code, Input: input}
}

func (e *Execution) TempFile() (*os.File, error) {
	var file *os.File
	var err error

	if e.Language == "java" {
		file, err = os.CreateTemp("./"+e.Language+"/", e.CodeID+"."+e.Language)
	} else {
		file, err = os.CreateTemp("./"+e.Language+"/", "Main.java")
	}

	defer file.Close()

	if err != nil {
		return nil, err
	}

	_, err = file.WriteString(e.Code)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (e *Execution) ExecuteCompilable(file *os.File) (*exec.Cmd, *models.Response, error) {
	output := new(models.Response)
	var stderr bytes.Buffer

	var cmd *exec.Cmd

	switch e.Language {
	case "c":
		cmd = exec.Command("gcc", file.Name(), "-o", "./"+e.Language+"/", e.CodeID)
	case "cpp":
		cmd = exec.Command("g++", file.Name(), "-o", "./"+e.Language+"/", e.CodeID)
	case "java":
		cmd = exec.Command("javac", "./"+e.Language+"/", file.Name())
	default:
		return nil, nil, errors.New("language not supported")
	}

	cmd.Stderr = &stderr
	output.Errors = stderr.String()

	// err
	err := cmd.Run()
	if err != nil {
		return nil, output, err
	}

	if e.Language == "java" {
		cmd = exec.Command("java", "./"+e.Language+"/", "Main")
	} else {
		cmd = exec.Command("./"+e.Language+"/", e.CodeID)
	}

	return cmd, nil, nil
}

func (e *Execution) ExecuteInterpreted(file *os.File) (*exec.Cmd, error) {
	var cmd *exec.Cmd

	switch e.Language {
	case "python3":
		cmd = exec.Command("python", "./"+"py"+"/", file.Name())
	case "go":
		cmd = exec.Command("go run", "./"+e.Language+"/", file.Name())
	case "bash":
		file.Chmod(755)
		cmd = exec.Command("./"+"sh"+"/", file.Name())
	default:
		return nil, errors.New("language not supported")
	}

	return cmd, nil
}

func (e *Execution) Run(cmd *exec.Cmd, file *os.File) (*models.Response, error) {
	output := new(models.Response)
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	if e.Input != "" {
		cmd.Stdin = bytes.NewBufferString(e.Input)
	}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return output, err
	}

	output.Errors = stderr.String()
	output.Std_output[0] = stdout.String()
	output.ReqID = e.CodeID

	err = os.Remove(file.Name())
	if err != nil {
		return output, err
	}

	return output, nil
}
