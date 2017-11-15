package model

// RuleInfo model
type RuleInfo struct {
	Description string                 `json:"description,omitempty"`
	Enabled     bool                   `json:"enabled,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Rule        map[string]interface{} `json:"rule,omitempty"`
	Scope       string                 `json:"scope,omitempty"`
}
