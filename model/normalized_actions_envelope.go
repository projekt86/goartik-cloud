package model

// NormalizedActionsEnvelope model
type NormalizedActionsEnvelope struct {
	Data      *[]NormalizedAction `json:"data,omitempty"`
	Error     *ResultError        `json:"error,omitempty"`
	DDIDS     string              `json:"ddids,omitempty"`
	DDID      string              `json:"ddid,omitempty"`
	UserID    string              `json:"uid,omitempty"`
	StartDate int64               `json:"startDate,omitempty"`
	EndDate   int64               `json:"endDate,omitempty"`
	Order     string              `json:"order,omitempty"`
	Next      string              `json:"next,omitempty"`
	Prev      string              `json:"prev,omitempty"`
	Count     int                 `json:"count,omitempty"`
	Size      int                 `json:"size,omitempty"`
}
