package model

// Message model
type Message struct {
	SDID      string                 `json:"sdid,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Timestamp int64                  `json:"ts,omitempty"`
}
