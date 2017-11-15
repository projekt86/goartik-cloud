package model

// ManifestVersionsEnvelope model
type ManifestVersionsEnvelope struct {
	Data  *ManifestVersions `json:"data,omitempty"`
	Error *ResultError      `json:"error,omitempty"`
}
