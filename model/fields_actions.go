package model

// FieldsActions model
type FieldsActions struct {
	Fields  map[string]interface{} `json:"fields,omitempty"`
	Actions map[string]interface{} `json:"actions,omitempty"`
}
