package model

// DeviceTokenEnvelope model
type DeviceTokenEnvelope struct {
	Data  *DeviceToken `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
