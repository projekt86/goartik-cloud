package model

// RulesEnvelope model
type RulesEnvelope struct {
	Data   *[]OutputRule `json:"data,omitempty"`
	Error  *ResultError  `json:"error,omitempty"`
	Total  int           `json:"total,omitempty"`
	Offset int           `json:"offset,omitempty"`
	Count  int           `json:"count,omitempty"`
}
