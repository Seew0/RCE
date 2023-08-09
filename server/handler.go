package server

import (
	"fmt"
	"rce/models"
	"rce/utils"

	"github.com/gin-gonic/gin"
)

var s Server

func codeHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var req models.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(404, gin.H{"error": "failed to bind"})
	}
	fmt.Printf("THE DATA IS HERE %v", req)
	utils.CodeRunner(req, s.OutputChan)
	fmt.Println("hey")

	fmt.Printf("The data is in output channel\n%v",<-s.OutputChan)

	resp := <-s.OutputChan
	if resp.ReqID == req.ReqID {
		c.JSON(200, gin.H{
			"output": resp,
		})
		return
	}
	c.JSON(500, gin.H{"error": "error occured while sending output"})
}
