package utils

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"rce/models"
)

type Execution struct {
	CodeID   string
	Language string
	Code     string
	Input    string
}

func ExecutionAdapter(codeid string, Language string, code string, input string) *Execution {
	return &Execution{CodeID: codeid, Language: Language, Code: code, Input: input}
}

func (e *Execution) TempFile() (*os.File, error) {
	var file *os.File
	var err error

	fmt.Println(e.Language)

	// Generate a unique filename for the temporary file using a random string
	// or use the provided CodeID for languages other than Java.
	filename := e.CodeID
	extensionType := e.Language
	if e.Language == "java" {
		filename = "Main"
	}
	if e.Language == "python3" {
		extensionType = "py"
	}
	if e.Language == "javascript"{
		extensionType = "js"
	}

	// Create the directory if it doesn't exist
	err = os.MkdirAll("./"+extensionType, os.ModePerm)
	if err != nil {
		log.Println("Error creating directory:", err)
		return nil, err
	}

	file, err = os.Create(filepath.Join("./"+extensionType, filename+"."+extensionType))
	if err != nil {
		log.Println("Error creating temp file:", err)
		return nil, err
	}

	// Close the file when the function returns to ensure it gets cleaned up properly.
	// Use defer after checking for errors, so the file is closed only if it was successfully created.
	defer file.Close()

	_, err = file.WriteString(e.Code)
	if err != nil {
		log.Println("Error writing to temp file:", err)
		return nil, err
	}

	// Sync the file to ensure data is written to disk before returning.
	err = file.Sync()
	if err != nil {
		log.Println("Error syncing temp file:", err)
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
		log.Println("err aaya hai running mein ",err)
		return output, err
	}

	output.Errors = stderr.String()
	output.Std_output[0] = stdout.String()
	output.ReqID = e.CodeID

	err = os.Remove(file.Name())
	if err != nil {
		log.Println("err aaya hai deleting mein ",err)
		return output, err
	}

	return output, nil
}
