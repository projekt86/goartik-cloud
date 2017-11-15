package model

// Subscription model
type Subscription struct {
	ID                   string `json:"id,omitempty"`
	AppID                string `json:"aid,omitempty"`
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
	CreatedOn            int64  `json:"createdOn,omitempty"`
	ModifiedOn           int64  `json:"modifiedOn,omitempty"`
	Status               string `json:"status,omitempty"`
}
