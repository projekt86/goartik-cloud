package model

// TrialsEnvelope model
type TrialsEnvelope struct {
	Data   *TrialArray  `json:"data,omitempty"`
	Error  *ResultError `json:"error,omitempty"`
	Total  int          `json:"total,omitempty"`
	Offset int          `json:"offset,omitempty"`
	Count  int          `json:"count,omitempty"`
}
