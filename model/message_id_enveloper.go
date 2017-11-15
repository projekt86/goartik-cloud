package model

// MessageIDEnvelope model
type MessageIDEnvelope struct {
	Data  *MessageID   `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
