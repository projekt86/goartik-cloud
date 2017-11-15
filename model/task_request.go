package model

// TaskRequest model
type TaskRequest struct {
	Filter         string         `json:"filter,omitempty"`
	TaskType       string         `json:"taskType,omitempty"`
	DeviceTypeID   string         `json:"dtid,omitempty"`
	Property       string         `json:"property,omitempty"`
	DeviceIDs      string         `json:"dids,omitempty"`
	TaskParameters TaskParameters `json:"taskParameters,omitempty"`
}
