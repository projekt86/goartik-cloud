package model

// NotifMessage model
type NotifMessage struct {
	UserID             string                 `json:"uid,omitempty"`
	CTS                int                    `json:"cts,omitempty"`
	Data               map[string]interface{} `json:"data,omitempty"`
	MID                string                 `json:"mid,omitempty"`
	SourceDeviceID     string                 `json:"sdid,omitempty"`
	SourceDeviceTypeID string                 `json:"sdtid,omitempty"`
	MV                 string                 `json:"mv,omitempty"`
	Timestamp          int64                  `json:"ts,omitempty"`
}
