package model

// DeviceRegStatusResponseEnvelope model
type DeviceRegStatusResponseEnvelope struct {
	Data  *DeviceRegStatusResponse `json:"data,omitempty"`
	Error *ResultError             `json:"error,omitempty"`
}
