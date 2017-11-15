package model

// UnregisterDeviceResponseEnvelope model
type UnregisterDeviceResponseEnvelope struct {
	Data  *UnregisterDeviceResponse `json:"data,omitempty"`
	Error *ResultError              `json:"error,omitempty"`
}
