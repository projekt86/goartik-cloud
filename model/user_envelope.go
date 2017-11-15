package model

// UserEnvelope model
type UserEnvelope struct {
	Data  *User        `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
