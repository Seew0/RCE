package server

import (
	"rce/models"
	"rce/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port       string
	router     *gin.Engine
	InputChan  chan models.Request
	OutputChan chan models.Response
}

func NewServer(port string, InputChan chan models.Request, OutputChan chan models.Response, router *gin.Engine) *Server {
	return &Server{
		port:       port,
		InputChan:  InputChan,
		OutputChan: OutputChan,
		router:     router,
	}
}

func (s *Server) Run() {
	go utils.CodeRunner(s.InputChan, s.OutputChan)

	s.router.POST("/api/execute", func(ctx *gin.Context) {
		codeHandler(ctx)
	})
}
