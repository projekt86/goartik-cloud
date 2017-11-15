package model

// RuleError model
type RuleError struct {
	ErrorCode   int       `json:"errorCode,omitempty"`
	FieldPath   FieldPath `json:"fieldPath,omitempty"`
	MessageArgs []string  `json:"messageArgs,omitempty"`
	MessageKey  string    `json:"messageKey,omitempty"`
}
