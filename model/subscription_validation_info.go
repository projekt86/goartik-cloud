package model

// SubscriptionValidationInfo model
type SubscriptionValidationInfo struct {
	AppID string `json:"aid,omitempty"`
	Nonce string `json:"nonce,omitempty"`
}
