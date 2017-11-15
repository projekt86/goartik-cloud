package model

// DeviceRegConfirmUserResponseEnvelope model
type DeviceRegConfirmUserResponseEnvelope struct {
	Data  *DeviceRegConfirmUserResponse `json:"data,omitempty"`
	Error *ResultError                  `json:"error,omitempty"`
}
