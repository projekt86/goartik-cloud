package model

// DeviceTypeInfo model
type DeviceTypeInfo struct {
	DevicePropertiesEnabled bool   `json:"devicePropertiesEnabled,omitempty"`
	PMax                    int    `json:"pmax,omitempty"`
	ModifiedOn              int64  `json:"modifiedOn,omitempty"`
	DeviceTypeID            string `json:"dtid,omitempty"`
	PMin                    int    `json:"pmin,omitempty"`
	TaskExpiresAfter        int    `json:"taskExpiresAfter,omitempty"`
	CreatedOn               int64  `json:"createdOn,omitempty"`
}
