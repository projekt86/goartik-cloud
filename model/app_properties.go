package model

// AppProperties model
type AppProperties struct {
	UserID     string `json:"uid,omitempty"`
	AppID      string `json:"aid,omitempty"`
	Properties string `json:"properties,omitempty"`
}
