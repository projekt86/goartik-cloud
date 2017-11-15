package model

// ExportHistoryEnvelope model
type ExportHistoryEnvelope struct {
	Data   *ExportDataArray `json:"data,omitempty"`
	Error  *ResultError     `json:"error,omitempty"`
	Count  int              `json:"count,omitempty"`
	Offset int              `json:"offset,omitempty"`
	Total  int              `json:"total,omitempty"`
}
