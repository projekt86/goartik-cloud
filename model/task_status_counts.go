package model

// TaskStatusCounts model
type TaskStatusCounts struct {
	TotalDevices int `json:"totalDevices,omitempty"`
	NumCompleted int `json:"numCompleted,omitempty"`
	NumSucceeded int `json:"numSucceeded,omitempty"`
	NumFailed    int `json:"numFailed,omitempty"`
	NumCancelled int `json:"numCancelled,omitempty"`
}
