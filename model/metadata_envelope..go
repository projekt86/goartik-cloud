package model

// MetadataEnvelope model
type MetadataEnvelope struct {
	Data  *map[string]interface{} `json:"data,omitempty"`
	Error *ResultError            `json:"error,omitempty"`
}
