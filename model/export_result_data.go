package model

// ExportResultData model
type ExportResultData struct {
	Data      map[string]int64 `json:"data,omitempty"`
	Timestamp int64            `json:"ts,omitempty"`
	MID       string           `json:"mid,omitempty"`
	SDTID     string           `json:"sdtid,omitempty"`
	CTS       int64            `json:"cts,omitempty"`
	UserID    string           `json:"uid,omitempty"`
	MV        int              `json:"mv,omitempty"`
	SDID      string           `json:"sdid,omitempty"`
}
