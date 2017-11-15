package model

// ExportResultEnvelope model
type ExportResultEnvelope struct {
	Data  *ExportResultData `json:"data,omitempty"`
	Error *ResultError      `json:"error,omitempty"`
	Size  int               `json:"size,omitempty"`
}
