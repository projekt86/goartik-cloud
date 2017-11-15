package model

// FieldPath model
type FieldPath struct {
	Path []NonEmptyString `json:"path,omitempty"`
}
