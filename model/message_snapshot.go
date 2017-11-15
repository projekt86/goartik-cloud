package model

// MessageSnapshot model
type MessageSnapshot struct {
	Data map[string]interface{} `json:"data,omitempty"`
	SDID string                 `json:"sdid,omitempty"`
}
