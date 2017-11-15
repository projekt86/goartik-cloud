package model

// TaskHistory model
type TaskHistory struct {
	NumAttempts  int    `json:"numAttempts,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	ErrorCode    string `json:"errorCode,omitempty"`
	DeviceID     string `json:"did,omitempty"`
	Status       string `json:"status,omitempty"`
	Timestamp    int64  `json:"ts,omitempty"`
}
