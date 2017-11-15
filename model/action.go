package model

// Action model
type Action struct {
	Name       string                 `json:"name,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}
