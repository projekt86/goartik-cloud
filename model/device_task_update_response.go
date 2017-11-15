package model

// DeviceTaskUpdateResponse model
type DeviceTaskUpdateResponse struct {
	Data  *TaskStatus  `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
