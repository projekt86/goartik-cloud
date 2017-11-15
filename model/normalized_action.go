package model

// NormalizedAction model
type NormalizedAction struct {
	Data      ActionArray `json:"data,omitempty"`
	DDTID     string      `json:"ddtid,omitempty"`
	SDID      string      `json:"sdid,omitempty"`
	DDID      string      `json:"ddid,omitempty"`
	UserID    string      `json:"uid,omitempty"`
	MessageID string      `json:"mid,omitempty"`
	Timestamp int64       `json:"ts,omitempty"`
	MV        int         `json:"mv,omitempty"`
	CTS       int         `json:"cts,omitempty"`
}
