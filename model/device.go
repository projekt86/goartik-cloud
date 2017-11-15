package model

// Device model
type Device struct {
	ID                    string                 `json:"id,omitempty"`
	UserUD                string                 `json:"uid,omitempty"`
	TypeID                string                 `json:"dtid,omitempty"`
	Name                  string                 `json:"name,omitempty"`
	ManifestVersion       int                    `json:"manifestVersion,omitempty"`
	ManifestVersionPolicy string                 `json:"manifestVersionPolicy,omitempty"`
	NeedProviderAuth      bool                   `json:"needProviderAuth,omitempty"`
	Properties            map[string]interface{} `json:"properties,omitempty"`
	CreatedOn             int                    `json:"createdOn,omitempty"`
	Connected             bool                   `json:"connected,omitempty"`
	CertificateInfo       string                 `json:"certificateInfo,omitempty"`
	CertificateSignature  string                 `json:"certificateSignature,omitempty"`
	ExternalID            string                 `json:"eid,omitempty"`
	ProviderCredentials   map[string]interface{} `json:"providerCredentials,omitempty"`
	SharedWithMe          string                 `json:"sharedWithMe,omitempty"`
	SharedWithOthers      bool                   `json:"sharedWithOthers,omitempty"`
}
