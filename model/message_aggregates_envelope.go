package model

// MessageAggregatesEnvelope model
type MessageAggregatesEnvelope struct {
	Data      *[]MessageAggregate `json:"data,omitempty"`
	Error     *ResultError        `json:"error,omitempty"`
	SDID      string              `json:"sdid,omitempty"`
	Field     string              `json:"field,omitempty"`
	StartDate int64               `json:"startDate,omitempty"`
	EndDate   int64               `json:"endDate,omitempty"`
	Size      int                 `json:"size,omitempty"`
}
