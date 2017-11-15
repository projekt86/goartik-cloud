package model

// TaskUpdateResponse model
type TaskUpdateResponse struct {
	Data  *Task        `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
