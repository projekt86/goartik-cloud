package model

// Presence model
type Presence struct {
	Connected  bool `json:"connected,omitempty"`
	LastSeenOn int  `json:"lastSeenOn,omitempty"`
}
