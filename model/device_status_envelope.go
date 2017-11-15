package model

// DeviceStatusEnvelope model
type DeviceStatusEnvelope struct {
	Data  *DeviceStatus `json:"data,omitempty"`
	Error *ResultError  `json:"error,omitempty"`
	DID   string        `json:"did,omitempty"`
}
