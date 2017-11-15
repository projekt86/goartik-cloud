package model

// PropertiesEnvelope model
type PropertiesEnvelope struct {
	Data  *AppProperties `json:"data,omitempty"`
	Error *ResultError   `json:"error,omitempty"`
}
