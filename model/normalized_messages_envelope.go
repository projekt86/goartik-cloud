package model

// NormalizedMessagesEnvelope model
type NormalizedMessagesEnvelope struct {
	Data      *[]NormalizedMessage `json:"data,omitempty"`
	Error     *ResultError         `json:"error,omitempty"`
	SDIDS     string               `json:"sdids,omitempty"`
	SDID      string               `json:"sdid,omitempty"`
	UserID    string               `json:"uid,omitempty"`
	StartDate int64                `json:"startDate,omitempty"`
	EndDate   int64                `json:"endDate,omitempty"`
	Order     string               `json:"order,omitempty"`
	Next      string               `json:"next,omitempty"`
	Prev      string               `json:"prev,omitempty"`
	Count     int                  `json:"count,omitempty"`
	Size      int                  `json:"size,omitempty"`
}
