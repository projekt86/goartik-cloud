package model

// DeviceStatus model
type DeviceStatus struct {
	LastMessageTS  int       `json:"lastMessageTs,omitempty"`
	LastActionTS   int       `json:"lastActionTs,omitempty"`
	Availability   string    `json:"availability,omitempty"`
	LastTimeOnline int       `json:"lastTimeOnline,omitempty"`
	Snapshot       *Snapshot `json:"snapshot,omitempty"`
}
