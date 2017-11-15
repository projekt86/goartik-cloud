package model

// DevicesStatusEnvelope model
type DevicesStatusEnvelope struct {
	Data   *[]DeviceStatusEnvelope `json:"data,omitempty"`
	Error  *ResultError            `json:"error,omitempty"`
	Total  int                     `json:"total,omitempty"`
	Offset int                     `json:"offset,omitempty"`
	Count  int                     `json:"count,omitempty"`
}
