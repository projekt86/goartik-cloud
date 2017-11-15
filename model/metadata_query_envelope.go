package model

// MetadataQueryEnvelope model
type MetadataQueryEnvelope struct {
	Data   *map[string]interface{} `json:"data,omitempty"`
	Error  *ResultError            `json:"error,omitempty"`
	Total  int                     `json:"total,omitempty"`
	Count  int                     `json:"count,omitempty"`
	Offset int                     `json:"offset,omitempty"`
}
