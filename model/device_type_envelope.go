package model

// DeviceTypeEnvelope model
type DeviceTypeEnvelope struct {
	Data  *DeviceType  `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
