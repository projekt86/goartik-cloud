package model

// TokenInfoEnvelope model
type TokenInfoEnvelope struct {
	Data  *TokenInfo   `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
