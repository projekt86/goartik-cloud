package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// SubscriptionsAPI model.
type SubscriptionsAPI struct {
	Client *client.APIClient
}

// NewSubscriptionsAPI returns new desubscriptionsvice API with default API client.
func NewSubscriptionsAPI() *SubscriptionsAPI {
	return &SubscriptionsAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewSubscriptionsAPIWithClient returns new subscriptions API with custom API client.
func NewSubscriptionsAPIWithClient(client *client.APIClient) *SubscriptionsAPI {
	return &SubscriptionsAPI{
		Client: client,
	}
}

// CreateSubscription creates a new subscription.
func (a *SubscriptionsAPI) CreateSubscription(subscriptionInfo *model.SubscriptionInfo) (*model.SubscriptionEnvelope, error) {
	if subscriptionInfo == nil {
		return nil, fmt.Errorf("subscriptionInfo must be provided")
	}

	u := fmt.Sprintf("%s/subscriptions", apiREST)

	b, err := a.Client.NewPostRequest(u, subscriptionInfo)
	if err != nil {
		return nil, err
	}

	result := &model.SubscriptionEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteSubscription deletes a subscription.
func (a *SubscriptionsAPI) DeleteSubscription(subscriptionID string) (*model.SubscriptionEnvelope, error) {
	u := fmt.Sprintf("%s/subscriptions/%s", apiREST, subscriptionID)

	b, err := a.Client.NewDeleteRequest(u)
	if err != nil {
		return nil, err
	}

	result := &model.SubscriptionEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAllSubscriptions returns all subscriptions for the current application.
func (a *SubscriptionsAPI) GetAllSubscriptions(userID string, count int, offset int) (*model.SubscriptionsEnvelope, error) {
	u := fmt.Sprintf("%s/subscriptions", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 1, math.MaxInt32))
	}
	v.Add("uid", userID)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.SubscriptionsEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetSubscription returns a subscription.
func (a *SubscriptionsAPI) GetSubscription(subscriptionID string) (*model.SubscriptionEnvelope, error) {
	u := fmt.Sprintf("%s/subscriptions/%s", apiREST, subscriptionID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.SubscriptionEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ValidateSubscription validate a subscription.
func (a *SubscriptionsAPI) ValidateSubscription(subscriptionID string, info *model.SubscriptionValidationInfo) (*model.SubscriptionEnvelope, error) {
	if info == nil {
		return nil, fmt.Errorf("info must be provided")
	}

	u := fmt.Sprintf("%s/subscriptions/%s/validate", apiREST, subscriptionID)

	b, err := a.Client.NewPostRequest(u, info)
	if err != nil {
		return nil, err
	}

	result := &model.SubscriptionEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
