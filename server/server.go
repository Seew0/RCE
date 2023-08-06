package server

import (
	"log"
	"rce/models"

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
	s.router.POST("/api/execute", func(ctx *gin.Context) {
		codeHandler(ctx)
	})

	log.Printf("Your server is serving at 127.0.0.1%v", s.port)
	s.router.Run(s.port)
}
