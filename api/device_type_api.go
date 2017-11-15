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

// DeviceTypesAPI model.
type DeviceTypesAPI struct {
	Client *client.APIClient
}

// NewDeviceTypesAPI returns new device type API with default API client.
func NewDeviceTypesAPI() *DeviceTypesAPI {
	return &DeviceTypesAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewDeviceTypesAPIWithClient returns new device type API with custom API client.
func NewDeviceTypesAPIWithClient(client *client.APIClient) *DeviceTypesAPI {
	return &DeviceTypesAPI{
		Client: client,
	}
}

// GetDeviceType returns a device type.
func (a *DeviceTypesAPI) GetDeviceType(deviceTypeID string) (*model.DeviceTypeEnvelope, error) {
	u := fmt.Sprintf("%s/devicetypes/%s", apiREST, deviceTypeID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTypeEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeviceTypes searches for device types by name.
func (a *DeviceTypesAPI) GetDeviceTypes(name string, count int, offset int, tags []string) (*model.DeviceTypesEnvelope, error) {
	u := fmt.Sprintf("%s/devicetypes", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}
	if tags != nil && len(tags) > 0 {
		v.Add("tags", strings.Join(tags, ","))
	}

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTypesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAvailableManifestVersions returns a device type's available Manifest versions.
func (a *DeviceTypesAPI) GetAvailableManifestVersions(deviceTypeID string) (*model.ManifestVersionsEnvelope, error) {
	u := fmt.Sprintf("%s/devicetypes/%s/availablemanifestversions", apiREST, deviceTypeID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.ManifestVersionsEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetLatestManifestProperties return lates manifest properties.
func (a *DeviceTypesAPI) GetLatestManifestProperties(deviceTypeID string) (*model.ManifestPropertiesEnvelope, error) {
	u := fmt.Sprintf("%s/devicetypes/%s/manifests/latest/properties", apiREST, deviceTypeID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.ManifestPropertiesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetManifestProperties returns properties for a specific Manifest version of a device type.
func (a *DeviceTypesAPI) GetManifestProperties(deviceTypeID string, version int) (*model.ManifestPropertiesEnvelope, error) {
	u := fmt.Sprintf("%s/devicetypes/%s/manifests/%d/properties", apiREST, deviceTypeID, version)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.ManifestPropertiesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeviceTypesByApplication returns an Application's device types from the scope parameters.
func (a *DeviceTypesAPI) GetDeviceTypesByApplication(appID string, count int, offset int) (*model.DeviceTypesEnvelope, error) {
	u := fmt.Sprintf("%s/applications/%s/devicetypes", apiREST, appID)

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

	result := &model.DeviceTypesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
