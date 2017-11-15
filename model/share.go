package model

// Share model
type Share struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   string `json:"status,omitempty"`
	SharedOn int    `json:"sharedOn,omitempty"`
}
