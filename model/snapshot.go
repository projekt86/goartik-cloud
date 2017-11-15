package model

// Snapshot model
type Snapshot struct {
	Data map[string]interface{} `json:"data,omitempty"`
	SDID string                 `json:"sdid,omitempty"`
}
