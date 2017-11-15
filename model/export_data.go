package model

// ExportData model
type ExportData struct {
	ExpirationDate int64  `json:"expirationDate,omitempty"`
	ExportID       string `json:"exportId,omitempty"`
	FileSize       int    `json:"fileSize,omitempty"`
	Status         string `json:"status,omitempty"`
	MD5            string `json:"md5,omitempty"`
	TotalMessages  int    `json:"totalMessages,omitempty"`
}
