package api

// Constants
const (
	ManifestVersionPolicyLatest = "LATEST"
	ManifestVersionPolicyDevice = "DEVICE"

	DeviceStatusOnline  = "online"
	DeviceStatusOffline = "offline"
	DeviceStatusUnknown = "unknown"

	OrderASC  = "asc"
	OrderDESC = "desc"

	SubscriptionTypeHTTPCallback = "httpCallback"
	SubscriptionTypeAWSKinesis   = "awsKinesis"

	RuleScopeAllApp  = "allApplications"
	RuleScopeThisApp = "thisApplication"
)
