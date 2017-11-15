package model

// Tag model
type Tag struct {
	Name       string `json:"name,omitempty"`
	IsCategory bool   `json:"isCategory,omitempty"`
}
