package model

// TaskStatusesEnvelope model
type TaskStatusesEnvelope struct {
	Data   *TaskStatuses `json:"data,omitempty"`
	Error  *ResultError  `json:"error,omitempty"`
	Total  int           `json:"total,omitempty"`
	Count  int           `json:"count,omitempty"`
	Offset int           `json:"offset,omitempty"`
}
