package model

// TagsEnvelope model
type TagsEnvelope struct {
	Data  *TagArray    `json:"data,omitempty"`
	Error *ResultError `json:"error,omitempty"`
}
