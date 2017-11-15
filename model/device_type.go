package model

// DeviceType model
type DeviceType struct {
	ID                 string `json:"id,omitempty"`
	UniqueName         string `json:"uniqueName,omitempty"`
	LatestVersion      int    `json:"latestVersion,omitempty"`
	LastUpdated        int    `json:"lastUpdated,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	UserID             string `json:"uid,omitempty"`
	OrganizationID     string `json:"oid,omitempty"`
	HasCloudConnector  bool   `json:"hasCloudConnector,omitempty"`
	Approved           bool   `json:"approved,omitempty"`
	Published          bool   `json:"published,omitempty"`
	Protected          bool   `json:"protected,omitempty"`
	InStore            bool   `json:"inStore,omitempty"`
	OwnedByCurrentUser bool   `json:"ownedByCurrentUser,omitempty"`
	Tags               []Tag  `json:"tags,omitempty"`
	RSP                bool   `json:"rsp,omitempty"`
	IssuerDN           string `json:"issuerDn,omitempty"`
	VID                string `json:"vid,omitempty"`
}
