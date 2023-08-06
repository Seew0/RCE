package main

import (
	"log"
	"os"
	"rce/models"
	"rce/server"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	
	port := os.Getenv("port")
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		inputChan := make(chan models.Request, 1)
		outputChan := make(chan models.Response, 1)
		Router := gin.Default()
		server := server.NewServer(port, inputChan, outputChan, Router)
		
		server.Run()
	}()
	wg.Wait()
}
