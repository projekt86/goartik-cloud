package api

import (
	"encoding/json"
	"fmt"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// RegistrationsAPI : registrations API client
type RegistrationsAPI struct {
	Client *client.APIClient
}

// NewRegistrationsAPI : Return new registrations API with default API client
func NewRegistrationsAPI() *RegistrationsAPI {
	return &RegistrationsAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewRegistrationsAPIWithClient : Return new registrations API with custom API client
func NewRegistrationsAPIWithClient(client *client.APIClient) *RegistrationsAPI {
	return &RegistrationsAPI{
		Client: client,
	}
}

// ConfirmUser : confirm user
func (a *RegistrationsAPI) ConfirmUser(deviceID string, deviceName string, pin string) (*model.DeviceRegConfirmUserResponseEnvelope, error) {
	u := fmt.Sprintf("%s/devices/registrations/pin", apiREST)

	payload := struct {
		DeviceID   string `json:"deviceId,omitempty"`
		DeviceName string `json:"deviceName,omitempty"`
		Pin        string `json:"pin"`
	}{
		DeviceID:   deviceID,
		DeviceName: deviceName,
		Pin:        pin,
	}

	b, err := a.Client.NewPutRequest(u, payload)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceRegConfirmUserResponseEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetRequestStatusForUser : get request status for user
func (a *RegistrationsAPI) GetRequestStatusForUser(requestID string) (*model.DeviceRegStatusResponseEnvelope, error) {
	u := fmt.Sprintf("%s/devices/registrations/%s/status", apiREST, requestID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceRegStatusResponseEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UnregisterDevice : unregister device
func (a *RegistrationsAPI) UnregisterDevice(deviceID string) (*model.UnregisterDeviceResponseEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/registrations/", apiREST, deviceID)

	b, err := a.Client.NewDeleteRequest(u)
	if err != nil {
		return nil, err
	}

	result := &model.UnregisterDeviceResponseEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
