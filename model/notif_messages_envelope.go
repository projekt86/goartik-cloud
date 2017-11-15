package model

// NotifMessagesEnvelope model
type NotifMessagesEnvelope struct {
	Data   *NotifMessageArray `json:"data,omitempty"`
	Error  *ResultError       `json:"error,omitempty"`
	Total  int                `json:"total,omitempty"`
	Count  int                `json:"count,omitempty"`
	Offset int                `json:"offset,omitempty"`
	Order  string             `json:"order,omitempty"`
}
