package main

import (
	"log"
	"os"
	"rce/models"
	"rce/server"
	"sync"

	"github.com/joho/godotenv"
)

// what if we use goroutines and channels and bwrap to scale,
// make a small vertically scalled program and then in horizontal just use sandboxes?

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// start bwrap monitor
	var wg sync.WaitGroup
	port := os.Getenv("port")

	wg.Add(1)
	go func() {
		defer wg.Done()
		inputChan := make(chan models.Request)
		outputChan := make(chan models.Response)
		server := server.NewServer(port, inputChan, outputChan)
		server.Run()
	}()
	wg.Wait()
}
