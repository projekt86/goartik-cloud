package model

// DevicesEnvelope model
type DevicesEnvelope struct {
	Data   *DeviceArray `json:"data,omitempty"`
	Error  *ResultError `json:"error,omitempty"`
	Total  int          `json:"total,omitempty"`
	Offset int          `json:"offset,omitempty"`
	Count  int          `json:"count,omitempty"`
}
