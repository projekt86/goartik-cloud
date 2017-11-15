package model

// DeviceShareEnvelope model
type DeviceShareEnvelope struct {
	Data  *Share       `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
