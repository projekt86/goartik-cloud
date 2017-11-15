package model

// Task model
type Task struct {
	Filter         string           `json:"filter,omitempty"`
	TaskType       string           `json:"taskType,omitempty"`
	ModifiedOn     int64            `json:"modifiedOn,omitempty"`
	DeviceTypeID   string           `json:"dtid,omitempty"`
	StatusCounts   TaskStatusCounts `json:"statusCounts,omitempty"`
	Property       string           `json:"property,omitempty"`
	ID             string           `json:"id,omitempty"`
	DeviceIDs      string           `json:"dids,omitempty"`
	TaskParameters TaskParameters   `json:"taskParameters,omitempty"`
	CreatedOn      int64            `json:"createdOn,omitempty"`
	Status         string           `json:"status,omitempty"`
}
