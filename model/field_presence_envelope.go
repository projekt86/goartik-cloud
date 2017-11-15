package model

// FieldPresenceEnvelope model
type FieldPresenceEnvelope struct {
	Data          *[]FieldPresence `json:"data,omitempty"`
	Error         *ResultError     `json:"error,omitempty"`
	SDID          string           `json:"sdid,omitempty"`
	FieldPresence string           `json:"fieldPresence,omitempty"`
	Interval      string           `json:"interval,omitempty"`
	StartDate     int64            `json:"startDate,omitempty"`
	EndDate       int64            `json:"endDate,omitempty"`
	Size          int              `json:"size,omitempty"`
}
