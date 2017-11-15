package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// UsersAPI model.
type UsersAPI struct {
	Client *client.APIClient
}

// NewUsersAPI returns new users API with default API client.
func NewUsersAPI() *UsersAPI {
	return &UsersAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewUsersAPIWithClient returns new users API with custom API client.
func NewUsersAPIWithClient(client *client.APIClient) *UsersAPI {
	return &UsersAPI{
		Client: client,
	}
}

// GetSelf returns the current user's profile.
func (a *UsersAPI) GetSelf() (*model.UserEnvelope, error) {
	url := fmt.Sprintf("%s/users/self", apiREST)

	b, err := a.Client.NewGetRequest(url, nil)
	if err != nil {
		return nil, err
	}

	result := &model.UserEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserDevices returns a user's devices.
func (a *UsersAPI) GetUserDevices(userID string, count int, offset int, includeProperties bool, owner string, includeShareInfo bool) (*model.DevicesEnvelope, error) {
	u := fmt.Sprintf("%s/users/%s/devices", apiREST, userID)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}
	v.Add("includeProperties", parseBoolParam(includeProperties))
	if owner != "" {
		v.Add("owner", parseStringParam(owner, []string{"ALL", "ME", "SHARED_WITH_ME"}, "ALL"))
	}
	v.Add("includeShareInfo", parseBoolParam(includeShareInfo))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.DevicesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserDeviceTypes returns a user's device types.
func (a *UsersAPI) GetUserDeviceTypes(userID string, count int, offset int, includeShared bool) (*model.DeviceTypesEnvelope, error) {
	u := fmt.Sprintf("%s/users/%s/devicetypes", apiREST, userID)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}
	v.Add("includeShared", parseBoolParam(includeShared))

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

// GetUserProperties returns a user's application properties.
func (a *UsersAPI) GetUserProperties(userID string, appID string) (*model.PropertiesEnvelope, error) {
	u := fmt.Sprintf("%s/users/%s/properties", apiREST, userID)

	// parse params
	v := url.Values{}
	v.Add("aid", appID)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.PropertiesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateUserProperties creates a user's application properties.
func (a *UsersAPI) CreateUserProperties(userID string, appID string, properties *model.AppProperties) (*model.PropertiesEnvelope, error) {
	if properties == nil {
		return nil, fmt.Errorf("properties must be set")
	}

	u := fmt.Sprintf("%s/users/%s/properties", apiREST, userID)

	b, err := a.Client.NewPostRequest(u, properties)
	if err != nil {
		return nil, err
	}

	result := &model.PropertiesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteUserProperties deletes a user's application properties.
func (a *UsersAPI) DeleteUserProperties(userID string, appID string) (*model.PropertiesEnvelope, error) {
	u := fmt.Sprintf("%s/users/%s/properties", apiREST, userID)

	b, err := a.Client.NewDeleteRequest(u)
	if err != nil {
		return nil, err
	}

	result := &model.PropertiesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateUserProperties update a user's application properties.
func (a *UsersAPI) UpdateUserProperties(userID string, appID string, properties *model.AppProperties) (*model.PropertiesEnvelope, error) {
	if properties == nil {
		return nil, fmt.Errorf("properties must be set")
	}

	u := fmt.Sprintf("%s/users/%s/properties", apiREST, userID)

	b, err := a.Client.NewPutRequest(u, properties)
	if err != nil {
		return nil, err
	}

	result := &model.PropertiesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserRules returns a user's rules.
func (a *UsersAPI) GetUserRules(userID string, count int, offset int, excludeDisabled bool, scope string) (*model.RulesEnvelope, error) {
	u := fmt.Sprintf("%s/users/%s/rules", apiREST, userID)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}
	if scope != "" {
		v.Add("scope", parseStringParam(scope, []string{"allApplications", "thisApplication", "allApplications,thisApplication"}, "allApplications,thisApplication"))
	}
	v.Add("excludeDisabled", parseBoolParam(excludeDisabled))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.RulesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetUserTrials returns a user's trials.
func (a *UsersAPI) GetUserTrials(userID string, count int, offset int, role string) (*model.TrialsEnvelope, error) {
	u := fmt.Sprintf("%s/users/%s/trials", apiREST, userID)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}
	if role != "" {
		v.Add("role", parseStringParam(role, []string{"administrator", "participant"}, "administrator"))
	}

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(*b))
	result := &model.TrialsEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
