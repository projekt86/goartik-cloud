package model

// SubscriptionInfo model
type SubscriptionInfo struct {
	CallbackURL          string `json:"callbackUrl,omitempty"`
	DestinationDeviceID  string `json:"ddid,omitempty"`
	Description          string `json:"description,omitempty"`
	MessageType          string `json:"messageType,omitempty"`
	Name                 string `json:"name,omitempty"`
	SourceDeviceID       string `json:"sdid,omitempty"`
	SourceDeviceTypeID   string `json:"sdtid,omitempty"`
	UserID               string `json:"uid,omitempty"`
	SubscriptionType     string `json:"subscriptionType,omitempty"`
	AWSKey               string `json:"awsKey,omitempty"`
	AWSSecret            string `json:"awsSecret,omitempty"`
	AWSRegion            string `json:"awsRegion,omitempty"`
	AWSKinesisStreamName string `json:"awsKinesisStreamName,omitempty"`
	IncludeSharedDevices bool   `json:"includeSharedDevices,omitempty"`
}
