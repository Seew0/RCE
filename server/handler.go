package server

import (
	"rce/models"

	"github.com/gin-gonic/gin"
)

var s Server

func codeHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var req models.Request
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(404, gin.H{"error":"failed to bind"})
	}

	s.InputChan <- req

	for {
		if (<-s.OutputChan).ReqID == req.ReqID {
			c.JSON(200, gin.H{
				"output": <-s.OutputChan,
			})
			break
		}
	}
	c.JSON(500, gin.H{"error": "error occured while sending output"})
}
