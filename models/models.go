package models

type Request struct {
	ReqID string `json:"rID,omitempty"`
	Code      string `json:"code,omitempty"`
	CodeInput string `json:"input"`
	Language string `json:"language,omitempty"`
}

type Response struct {
	ReqID  string `json:"reqID,omitempty"`
	Output string `json:"output"`
	Errors string `json:"errors"`
}
