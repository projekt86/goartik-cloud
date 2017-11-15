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

// DevicesAPI model
type DevicesAPI struct {
	Client *client.APIClient
}

// NewDevicesAPI returns new device API with default API client
func NewDevicesAPI() *DevicesAPI {
	return &DevicesAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewDevicesAPIWithClient returns new device API with custom API client
func NewDevicesAPIWithClient(client *client.APIClient) *DevicesAPI {
	return &DevicesAPI{
		Client: client,
	}
}

// GetDevice returns a device
func (a *DevicesAPI) GetDevice(deviceID string, includeProperties bool) (*model.DeviceEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s", apiREST, deviceID)

	// parse params
	v := url.Values{}
	v.Add("properties", parseBoolParam(includeProperties))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AddDevice creates a device
func (a *DevicesAPI) AddDevice(device *model.Device) (*model.DeviceEnvelope, error) {
	if device == nil {
		return nil, fmt.Errorf("you must provice device")
	}

	u := fmt.Sprintf("%s/devices", apiREST)

	b, err := a.Client.NewPostRequest(u, device)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateDevice updates a device
func (a *DevicesAPI) UpdateDevice(device *model.Device) (*model.DeviceEnvelope, error) {
	if device == nil {
		return nil, fmt.Errorf("you must provice device")
	}

	u := fmt.Sprintf("%s/devices/%s", apiREST, device.ID)

	b, err := a.Client.NewPutRequest(u, device)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteDevice deletes a device
func (a *DevicesAPI) DeleteDevice(deviceID string) (*model.DeviceEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s", apiREST, deviceID)

	b, err := a.Client.NewDeleteRequest(u)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeviceToken returns a device's token
func (a *DevicesAPI) GetDeviceToken(deviceID string) (*model.DeviceTokenEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/tokens", apiREST, deviceID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTokenEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateDeviceToken creates a device's token
func (a *DevicesAPI) CreateDeviceToken(deviceID string) (*model.DeviceTokenEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/tokens", apiREST, deviceID)

	b, err := a.Client.NewPutRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTokenEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteDeviceToken deletes a device's token
func (a *DevicesAPI) DeleteDeviceToken(deviceID string) (*model.DeviceTokenEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/tokens", apiREST, deviceID)

	b, err := a.Client.NewDeleteRequest(u)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTokenEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDevicePresence returns the presence status of the device along with the time it was last seen
func (a *DevicesAPI) GetDevicePresence(deviceID string) (*model.PresenceEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/presence", apiREST, deviceID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.PresenceEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeviceStatus returns a device status
func (a *DevicesAPI) GetDeviceStatus(deviceID string, includeSnapshotTimestamp bool, includeSnapshot bool) (*model.DeviceStatusEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/status", apiREST, deviceID)

	// parse params
	v := url.Values{}
	v.Add("includeSnapshotTimestamp", parseBoolParam(includeSnapshotTimestamp))
	v.Add("includeSnapshot", parseBoolParam(includeSnapshot))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceStatusEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDevicesStatus returns devices status
func (a *DevicesAPI) GetDevicesStatus(deviceIDs []string, includeSnapshotTimestamp bool, includeSnapshot bool) (*model.DevicesStatusEnvelope, error) {
	if len(deviceIDs) == 0 {
		return nil, fmt.Errorf("you must provice id list")
	}

	u := fmt.Sprintf("%s/devices/status", apiREST)

	// parse params
	v := url.Values{}
	v.Add("includeSnapshotTimestamp", parseBoolParam(includeSnapshotTimestamp))
	v.Add("includeSnapshot", parseBoolParam(includeSnapshot))
	v.Add("dids", strings.Join(deviceIDs, ","))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.DevicesStatusEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateDeviceStatus updates a device status
func (a *DevicesAPI) UpdateDeviceStatus(deviceID string, device *model.DeviceStatusEnvelope) (*model.DeviceStatusEnvelope, error) {
	if device == nil {
		return nil, fmt.Errorf("you must provice device")
	}

	u := fmt.Sprintf("%s/devices/%s/status", apiREST, deviceID)

	b, err := a.Client.NewPutRequest(u, device)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceStatusEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateDeviceShare creates device share
func (a *DevicesAPI) CreateDeviceShare(deviceID string, email string) (*model.DeviceSharesEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/shares", apiREST, deviceID)

	payload := struct {
		Email string `json:"email"`
	}{
		Email: email,
	}

	b, err := a.Client.NewPostRequest(u, payload)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceSharesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeviceShares returns list all shares for the given device id
func (a *DevicesAPI) GetDeviceShares(deviceID string, count int, offset int) (*model.DeviceSharesEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/shares", apiREST, deviceID)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceSharesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeviceShare returns specific share of the given device id
func (a *DevicesAPI) GetDeviceShare(deviceID string, shareID string) (*model.DeviceShareEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/shares/%s", apiREST, deviceID, shareID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceShareEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteDeviceShare deletes device share
func (a *DevicesAPI) DeleteDeviceShare(deviceID string, shareID string) (*model.DeviceShareEnvelope, error) {
	u := fmt.Sprintf("%s/devices/%s/shares/%s", apiREST, deviceID, shareID)

	b, err := a.Client.NewDeleteRequest(u)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceShareEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
