package model

// ExportRequestData model
type ExportRequestData struct {
	CSVHeaders          bool   `json:"csvHeaders,omitempty"`
	EndDate             int64  `json:"endDate,omitempty"`
	ExportID            string `json:"exportId,omitempty"`
	Format              string `json:"format,omitempty"`
	Order               string `json:"order,omitempty"`
	SourceDeviceIDs     string `json:"sdids,omitempty"`
	SourceDeviceTypeIDs string `json:"sdtids,omitempty"`
	StartDate           int64  `json:"startDate,omitempty"`
	UserIDs             string `json:"uids,omitempty"`
	URL                 string `json:"url,omitempty"`
	TrailID             string `json:"trialId,omitempty"`
}
