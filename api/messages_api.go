package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"strings"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// MessagesAPI model.
type MessagesAPI struct {
	Client *client.APIClient
}

// NewMessagesAPI returns new messages API with default API client.
func NewMessagesAPI() *MessagesAPI {
	return &MessagesAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewMessagesAPIWithClient returns messages API with custom API client.
func NewMessagesAPIWithClient(client *client.APIClient) *MessagesAPI {
	return &MessagesAPI{
		Client: client,
	}
}

// SendMessage sends a Message from a Source Device.
func (a *MessagesAPI) SendMessage(message *model.Message) (*model.MessageIDEnvelope, error) {
	if message == nil {
		return nil, fmt.Errorf("message must be provided")
	}

	u := fmt.Sprintf("%s/messages", apiREST)

	b, err := a.Client.NewPostRequest(u, message)
	if err != nil {
		return nil, err
	}

	result := &model.MessageIDEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// SendActions sends ctions to a Destination Device
func (a *MessagesAPI) SendActions(actions *model.Actions) (*model.MessageIDEnvelope, error) {
	if actions == nil {
		return nil, fmt.Errorf("message must be provided")
	}

	u := fmt.Sprintf("%s/actions", apiREST)

	b, err := a.Client.NewPostRequest(u, actions)
	if err != nil {
		return nil, err
	}

	result := &model.MessageIDEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetNormalizedMessages returns normalized messages.
func (a *MessagesAPI) GetNormalizedMessages(userID string, sourceDeviceID string, messageID string, count int, offset int, startDate int64, endDate int64, order string, filter string, fieldPresence string) (*model.NormalizedMessagesEnvelope, error) {
	u := fmt.Sprintf("%s/messages", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 1, math.MaxInt32))
	}
	if startDate > 0 {
		v.Add("startDate", parseInt64Param(startDate, 1, math.MaxInt64))
	}
	if endDate > 0 {
		v.Add("endDate", parseInt64Param(endDate, 1, math.MaxInt64))
	}
	v.Add("uid", userID)
	v.Add("sdid", sourceDeviceID)
	v.Add("mid", messageID)
	v.Add("filter", filter)
	v.Add("fieldPresence", fieldPresence)
	v.Add("order", order)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.NormalizedMessagesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetNormalizedActions returns normalized actions.
func (a *MessagesAPI) GetNormalizedActions(userID string, destinationDeviceID string, messageID string, count int, offset int, startDate int64, endDate int64, order string) (*model.NormalizedActionsEnvelope, error) {
	u := fmt.Sprintf("%s/actions", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 1, math.MaxInt32))
	}
	if startDate > 0 {
		v.Add("startDate", parseInt64Param(startDate, 1, math.MaxInt64))
	}
	if endDate > 0 {
		v.Add("endDate", parseInt64Param(endDate, 1, math.MaxInt64))
	}
	v.Add("uid", userID)
	v.Add("mid", messageID)
	v.Add("order", order)
	v.Add("ddid", destinationDeviceID)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.NormalizedActionsEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetLastNormalizedMessages returns last normalized messages.
func (a *MessagesAPI) GetLastNormalizedMessages(count int, fieldPresence string, sourceDevicesIDs []string) (*model.NormalizedMessagesEnvelope, error) {
	u := fmt.Sprintf("%s/messages/last", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	v.Add("fieldPresence", fieldPresence)
	if sourceDevicesIDs != nil && len(sourceDevicesIDs) > 0 {
		v.Add("sdids", strings.Join(sourceDevicesIDs, ","))
	}

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.NormalizedMessagesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetNormalizedMessagesPresence returns normalized message presence.
// Returns bucketed timelines (in milliseconds since epoch) for which there was at least one message received.
// The grouping interval constrains the duration as:
//	If the interval is minute, you can query up to a maximum of one hour at a time
//	If the interval is hour, you can query up to a maximum of 24 hours at a time
//	If the interval is day, you can query up to a maximum of 31 days at a time
//	If the interval is month, you can query up to a maximum of 12 months at a time
//	If the interval is year, you can query up to a maximum of 10 years at a time
func (a *MessagesAPI) GetNormalizedMessagesPresence(sourceDevicesID string, interval string, fieldPresence string, startDate int64, endDate int64) (*model.FieldPresenceEnvelope, error) {
	u := fmt.Sprintf("%s/messages/presence", apiREST)

	// parse params
	v := url.Values{}
	if startDate > 0 {
		v.Add("startDate", parseInt64Param(startDate, 1, math.MaxInt64))
	}
	if endDate > 0 {
		v.Add("endDate", parseInt64Param(endDate, 1, math.MaxInt64))
	}
	v.Add("fieldPresence", fieldPresence)
	v.Add("sdid", sourceDevicesID)
	v.Add("interval", interval)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.FieldPresenceEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetNormalizedMessageAggregates returns aggregates on normalized messages.
func (a *MessagesAPI) GetNormalizedMessageAggregates(sourceDevicesID string, field string, startDate int64, endDate int64) (*model.MessageAggregatesEnvelope, error) {
	u := fmt.Sprintf("%s/messages/analytics/aggregates", apiREST)

	// parse params
	v := url.Values{}
	if startDate > 0 {
		v.Add("startDate", parseInt64Param(startDate, 1, math.MaxInt64))
	}
	if endDate > 0 {
		v.Add("endDate", parseInt64Param(endDate, 1, math.MaxInt64))
	}
	v.Add("sdid", sourceDevicesID)
	v.Add("field", field)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.MessageAggregatesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetNormalizedMessageHistogram returns histogram on normalized messages.
func (a *MessagesAPI) GetNormalizedMessageHistogram(sourceDeviceID string, field string, interval string, startDate int64, endDate int64) (*model.MessageAggregatesEnvelope, error) {
	u := fmt.Sprintf("%s/messages/analytics/histogram", apiREST)

	// parse params
	v := url.Values{}
	if startDate > 0 {
		v.Add("startDate", parseInt64Param(startDate, 1, math.MaxInt64))
	}
	if endDate > 0 {
		v.Add("endDate", parseInt64Param(endDate, 1, math.MaxInt64))
	}
	v.Add("sdid", sourceDeviceID)
	v.Add("field", field)
	v.Add("interval", interval)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.MessageAggregatesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetMessageSnapshots returns last received value for all Manifest fields (aka device state) of a device type.
func (a *MessagesAPI) GetMessageSnapshots(sourceDeviceIDs []string, includeTimestamp bool) (*model.MessageSnapshotsEnvelope, error) {
	if sourceDeviceIDs == nil || len(sourceDeviceIDs) == 0 || len(sourceDeviceIDs) > 100 {
		return nil, fmt.Errorf("list of source device IDs must be between 1 - 100")
	}

	u := fmt.Sprintf("%s/messages/snapshots", apiREST)

	// parse params
	v := url.Values{}
	v.Add("sdids", strings.Join(sourceDeviceIDs, ","))
	v.Add("includeTimestamp", parseBoolParam(includeTimestamp))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.MessageSnapshotsEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
