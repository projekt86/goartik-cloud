package model

// DeviceTasksEnvelope model
type DeviceTasksEnvelope struct {
	Data   *TaskArray   `json:"data,omitempty"`
	Error  *ResultError `json:"error,omitempty"`
	Total  int          `json:"total,omitempty"`
	Count  int          `json:"count,omitempty"`
	Offset int          `json:"offset,omitempty"`
}
