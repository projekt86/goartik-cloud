package model

// MessageHistogramEnvelope model
type MessageHistogramEnvelope struct {
	Data      *[]MessageHistogram `json:"data,omitempty"`
	Error     *ResultError        `json:"error,omitempty"`
	SDID      string              `json:"sdid,omitempty"`
	Field     string              `json:"field,omitempty"`
	StartDate int64               `json:"startDate,omitempty"`
	EndDate   int64               `json:"endDate,omitempty"`
	Size      int                 `json:"size,omitempty"`
	Interval  string              `json:"interval,omitempty"`
}
