package model

// ResultError model
type ResultError struct {
	Code    int         `json:"code,omitempty"`
	ID      string      `json:"id,omitempty"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}
