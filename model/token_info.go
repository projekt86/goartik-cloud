package model

// TokenInfo model
type TokenInfo struct {
	ClientID  string `json:"clientId,omitempty"`
	ExpiresIn int    `json:"expiresIn,omitempty"`
	UserID    string `json:"userId,omitempty"`
	DeviceID  string `json:"deviceId,omitempty"`
}
