package server

import (
	"fmt"
	"rce/models"
	"rce/utils"
	"sync"

	"github.com/gin-gonic/gin"
)

var s Server

var wg sync.WaitGroup

func codeHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var req models.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{"error": "failed to bind"})
	}
	fmt.Printf("THE DATA IS HERE %v", req)

	  go func() {
            for {
                s.InputChan <- req
            }
        }()

	utils.CodeRunner(s.InputChan, s.OutputChan)

	resp := <-s.OutputChan
	if resp.ReqID == req.ReqID {
		c.JSON(200, gin.H{
			"output": resp,
		})
		return
	}
	c.JSON(500, gin.H{"error": "error occured while sending output"})
}
