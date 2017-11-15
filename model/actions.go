package model

// Actions model
type Actions struct {
	DDID      string      `json:"ddid,omitempty"`
	Type      string      `json:"type,omitempty"`
	Data      ActionArray `json:"data,omitempty"`
	Timestamp int64       `json:"ts,omitempty"`
}
