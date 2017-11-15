package model

// SubscriptionEnvelope model
type SubscriptionEnvelope struct {
	Data  *Subscription `json:"data,omitempty"`
	Error *ResultError  `json:"error,omitempty"`
}
