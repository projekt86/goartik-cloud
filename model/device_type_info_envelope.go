package model

// DeviceTypeInfoEnvelope model
type DeviceTypeInfoEnvelope struct {
	Data  *DeviceTypeInfo `json:"data,omitempty"`
	Error *ResultError    `json:"error,omitempty"`
}
