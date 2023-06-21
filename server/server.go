package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"rce/models"
	"rce/utils"
)

type Server struct {
	port     string
	InputChan  chan models.Request
	OutputChan chan models.Response
}

func NewServer(port string, InputChan chan models.Request, OutputChan chan models.Response) *Server {
	return &Server{
		port:     port,
		InputChan:  InputChan,
		OutputChan: OutputChan,
	}
}

func (s *Server) Run() {
	go utils.CodeRunner(s.InputChan, s.OutputChan)

	http.HandleFunc("/execute", s.codeHandler)
	log.Fatal(http.ListenAndServe(s.port, nil))
}

func (s *Server) codeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.Request
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &req)

	s.InputChan <- req

	for {
		if (<-s.OutputChan).ReqID == req.ReqID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(<-s.OutputChan)
			break
		}
	}

	w.WriteHeader(http.StatusInternalServerError)
}
