package model

// RuleEnvelope model
type RuleEnvelope struct {
	Data  *OutputRule  `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
