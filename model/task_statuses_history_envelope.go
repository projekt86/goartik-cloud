package model

// TaskStatusesHistoryEnvelope model
type TaskStatusesHistoryEnvelope struct {
	Data  *TaskHistoryArray `json:"data,omitempty"`
	Error *ResultError      `json:"error,omitempty"`
}
