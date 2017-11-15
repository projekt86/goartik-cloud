package model

// PresenceEnvelope model
type PresenceEnvelope struct {
	Data  *Presence
	Error *ResultError
	SDID  string `json:"sdid"`
}
