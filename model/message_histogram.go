package model

// MessageHistogram model
type MessageHistogram struct {
	Count     int     `json:"count,omitempty"`
	Max       float32 `json:"max,omitempty"`
	Mean      float32 `json:"mean,omitempty"`
	Min       float32 `json:"min,omitempty"`
	Sum       float32 `json:"sum,omitempty"`
	Variance  float32 `json:"variance,omitempty"`
	Timestamp int64   `json:"ts,omitempty"`
}
