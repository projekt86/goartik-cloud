package model

// DeviceEnvelope model
type DeviceEnvelope struct {
	Data  *Device      `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
