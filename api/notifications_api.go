package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// NotificationsAPI model.
type NotificationsAPI struct {
	Client *client.APIClient
}

// NewNotificationsAPI returns new notifications API with default API client.
func NewNotificationsAPI() *NotificationsAPI {
	return &NotificationsAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewNotificationsAPIWithClient returns new notifications API with custom API client.
func NewNotificationsAPIWithClient(client *client.APIClient) *NotificationsAPI {
	return &NotificationsAPI{
		Client: client,
	}
}

// GetMessages returns messages associated with a notification.
func (a *NotificationsAPI) GetMessages(notifID string, count int, offset int, order string) (*model.NotifMessagesEnvelope, error) {
	u := fmt.Sprintf("%s/notifications/%s/messages", apiREST, notifID)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 1, math.MaxInt32))
	}
	v.Add("order", order)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.NotifMessagesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
