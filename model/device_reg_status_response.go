package model

// DeviceRegStatusResponse model
type DeviceRegStatusResponse struct {
	DeviceID string `json:"did,omitempty"`
	Status   string `json:"status,omitempty"`
}
