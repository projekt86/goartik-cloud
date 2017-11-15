package model

// MessageSnapshotsEnvelope model
type MessageSnapshotsEnvelope struct {
	Data  *[]MessageSnapshot `json:"data,omitempty"`
	Error *ResultError       `json:"error,omitempty"`
	SDIDS string             `json:"sdids,omitempty"`
	Size  int                `json:"size,omitempty"`
}
