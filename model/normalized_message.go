package model

// NormalizedMessage model
type NormalizedMessage struct {
	Data      map[string]interface{} `json:"data,omitempty"`
	SDTID     string                 `json:"sdtid,omitempty"`
	SDID      string                 `json:"sdid,omitempty"`
	UserID    string                 `json:"uid,omitempty"`
	MessageID string                 `json:"mid,omitempty"`
	Timestamp int64                  `json:"ts,omitempty"`
	MV        int                    `json:"mv,omitempty"`
	CTS       int                    `json:"cts,omitempty"`
}
