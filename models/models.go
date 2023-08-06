package models

type Request struct {
	ReqID string `json:"rID,omitempty"`
	Code      string `json:"code,omitempty"`
	CodeInput string `json:"input"`
	Language string `json:"language,omitempty"`
}

type Response struct{
	ReqID  string `json:"submission_id,omitempty"`
	Status_Code int `json:"status_code,omitempty"`
	Lang string `json:"lang,omitempty"`
	Status_runtime string `json:"status_runtime,omitempty"`
	// Memory int `json:"memory,omitempty"`
	Code_answer []string `json:"code_answer,omitempty"`
	Code_output []string `json:"code_output,omitempty"`
	Std_output_list []string `json:"std_output_list,omitempty"`
	Std_output []string `json:"std_output,omitempty"`
	// Elapsed_time int `json:"elapsed_time,omitempty"`
	// Task_finish_time int `json:"task_finish_time,omitempty"`
	// Task_name string `json:"task_name,omitempty"`
	// Total_correct string `json:"-"`
	// Total_testcase string `json:"-"`
	// Runtime_percentile string `json:"-"`
	// Status_memory string `json:"status_memory,omitempty"`
	// Memory_percentile string `json:"-"`
	// Pretty_lang string `json:"pretty_lang,omitempty"`
	// Status_msg string `json:"status_msg,omitempty"`
	// State string `json:"state,omitempty"`
	Errors string `json:"errors"`
}
