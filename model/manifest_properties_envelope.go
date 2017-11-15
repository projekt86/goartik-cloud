package model

// ManifestPropertiesEnvelope model
type ManifestPropertiesEnvelope struct {
	Data  *ManifestProperties `json:"data,omitempty"`
	Error *ResultError        `json:"error,omitempty"`
}
