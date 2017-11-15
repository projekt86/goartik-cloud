package model

// TaskEnvelope model
type TaskEnvelope struct {
	Data  *Task        `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
