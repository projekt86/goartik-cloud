package model

// ExportRequestEnvelope model
type ExportRequestEnvelope struct {
	Data  *ExportRequestData `json:"data,omitempty"`
	Error *ResultError       `json:"error,omitempty"`
}
