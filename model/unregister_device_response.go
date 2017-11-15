package model

// UnregisterDeviceResponse model
type UnregisterDeviceResponse struct {
	ID                    string `json:"id,omitempty"`
	UserID                string `json:"uid,omitempty"`
	ExternalID            string `json:"eid,omitempty"`
	DeviceID              string `json:"dtid,omitempty"`
	CertificateInfo       string `json:"certificateInfo,omitempty"`
	CertificateSignature  string `json:"certificateSignature,omitempty"`
	CreatedOn             int    `json:"createdOn,omitempty"`
	ManifestVersion       string `json:"manifestVersion,omitempty"`
	ManifestVersionPolicy string `json:"manifestVersionPolicy,omitempty"`
	Name                  string `json:"name,omitempty"`
	NeedProviderAuth      bool   `json:"needProviderAuth,omitempty"`
}
