package model

// DeviceToken model
type DeviceToken struct {
	UserUD      string `json:"uid,omitempty"`
	DeviceID    string `json:"did,omitempty"`
	AppID       string `json:"cid,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}
