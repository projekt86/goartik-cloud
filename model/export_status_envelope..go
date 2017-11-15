package model

// ExportStatusEnvelope model
type ExportStatusEnvelope struct {
	Data           *ExportData  `json:"data,omitempty"`
	Error          *ResultError `json:"error,omitempty"`
	ExpirationDate int64        `json:"expirationDate,omitempty"`
	ExportID       string       `json:"exportId,omitempty"`
	FileSize       int          `json:"fileSize,omitempty"`
	Status         string       `json:"status,omitempty"`
	MD5            string       `json:"md5,omitempty"`
	TotalMessages  int          `json:"totalMessages,omitempty"`
	TTL            string       `json:"ttl,omitempty"`
}
