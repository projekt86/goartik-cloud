package model

// TaskParameters model
type TaskParameters struct {
	ExpiresAfter int              `json:"expiresAfter,omitempty"`
	ScheduledOn  int              `json:"scheduledOn,omitempty"`
	Update       UpdateParameters `json:"update,omitempty"`
	Value        string           `json:"value,omitempty"`
}
