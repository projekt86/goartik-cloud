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

// DevicesManagementAPI model.
type DevicesManagementAPI struct {
	Client *client.APIClient
}

// NewDevicesManagementAPI returns new device management API with default API client.
func NewDevicesManagementAPI() *DevicesManagementAPI {
	return &DevicesManagementAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewDevicesManagementAPIWithClient returns new device management API with custom API client.
func NewDevicesManagementAPIWithClient(client *client.APIClient) *DevicesManagementAPI {
	return &DevicesManagementAPI{
		Client: client,
	}
}

// GetDeviceProperties a device's properties.
func (a *DevicesManagementAPI) GetDeviceProperties(deviceID string, includeTimestamp bool) (*model.MetadataEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/devices/%s/properties", apiREST, deviceID)

	// parse params
	v := url.Values{}
	v.Add("includeTimestamp", parseBoolParam(includeTimestamp))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.MetadataEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// QueryDeviceProperties query device properties across devices.
func (a *DevicesManagementAPI) QueryDeviceProperties(deviceTypeID string, count int, offset int, filter []string, includeTimestamp bool) (*model.MetadataQueryEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/devices/properties", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(count, 0, math.MaxInt32))
	}
	v.Add("includeTimestamp", parseBoolParam(includeTimestamp))
	v.Add("dtid", deviceTypeID)
	v.Add("filter", strings.Join(filter, ","))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.MetadataQueryEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateDeviceServerProperties updates a device's server properties.
func (a *DevicesManagementAPI) UpdateDeviceServerProperties(deviceID string, data *map[string]interface{}) (*model.MetadataEnvelope, error) {
	if data == nil {
		return nil, fmt.Errorf("data must be provided")
	}

	u := fmt.Sprintf("%s/devicemgmt/devices/%s/serverproperties", apiREST, deviceID)

	b, err := a.Client.NewPostRequest(u, data)
	if err != nil {
		return nil, err
	}

	result := &model.MetadataEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteDeviceServerProperties deletes a device's properties.
func (a *DevicesManagementAPI) DeleteDeviceServerProperties(deviceID string) (*model.MetadataEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/devices/%s/serverproperties", apiREST, deviceID)

	b, err := a.Client.NewDeleteRequest(u)
	if err != nil {
		return nil, err
	}

	result := &model.MetadataEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateTask updates a task for all devices - For now just allows changing the state to cancelled.
func (a *DevicesManagementAPI) UpdateTask(taskID string, status string) (*model.TaskUpdateResponse, error) {
	u := fmt.Sprintf("%s/devicemgmt/tasks/%s", apiREST, taskID)

	payload := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	b, err := a.Client.NewPutRequest(u, payload)
	if err != nil {
		return nil, err
	}

	result := &model.TaskUpdateResponse{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTaskByID returns the details and global status of a specific task id.
func (a *DevicesManagementAPI) GetTaskByID(taskID string) (*model.TaskEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/tasks/%s", apiREST, taskID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.TaskEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateTask create a new task for one or more devices.
func (a *DevicesManagementAPI) CreateTask(request *model.TaskRequest) (*model.TaskEnvelope, error) {
	if request == nil {
		return nil, fmt.Errorf("request must be provided")
	}

	u := fmt.Sprintf("%s/devicemgmt/tasks", apiREST)

	b, err := a.Client.NewPostRequest(u, request)
	if err != nil {
		return nil, err
	}

	result := &model.TaskEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTaskStatuses returns the details and status of a task id and the individual statuses of each device id in the list.
func (a *DevicesManagementAPI) GetTaskStatuses(taskID string, devicesIDs []string, status []string, count int, offset int) (*model.TaskStatusesEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/tasks/%s/statuses", apiREST, taskID)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}
	v.Add("dids", strings.Join(devicesIDs, ","))
	v.Add("status", strings.Join(status, ","))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.TaskStatusesEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTasks returns the all the tasks for a device type.
func (a *DevicesManagementAPI) GetTasks(deviceTypeID string, status []string, order string, sort string, count int, offset int) (*model.TasksEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/tasks", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}
	v.Add("dtid", deviceTypeID)
	v.Add("status", strings.Join(status, ","))
	v.Add("order", order)
	v.Add("sort", sort)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.TasksEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTaskStatusesHistory returns the history of the status changes for a specific task id, or for a specific device id in that task.
func (a *DevicesManagementAPI) GetTaskStatusesHistory(taskID string, deviceID string) (*model.TaskStatusesHistoryEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/tasks/%s/statuses/history", apiREST, taskID)

	// parse params
	v := url.Values{}
	v.Add("did", deviceID)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.TaskStatusesHistoryEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateTaskForDevice updates a task for a specific device - For now just allows changing the state to cancelled.
func (a *DevicesManagementAPI) UpdateTaskForDevice(taskID string, deviceID string, status string) (*model.DeviceTaskUpdateResponse, error) {
	u := fmt.Sprintf("%s/devicemgmt/tasks/%s/devices/%s", apiREST, taskID, deviceID)

	payload := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	b, err := a.Client.NewPutRequest(u, payload)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTaskUpdateResponse{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTasksByDevice returns the list of tasks for a particular device id with optional status filter.
func (a *DevicesManagementAPI) GetTasksByDevice(deviceID string, status []string, sort string, order string, count int, offset int) (*model.DeviceTasksEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/devices/%s/tasks", apiREST, deviceID)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 0, math.MaxInt32))
	}
	v.Add("status", strings.Join(status, ","))
	v.Add("sort", sort)
	v.Add("order", order)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTasksEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetManifestProperties returns a device type's device management manifest properties.
func (a *DevicesManagementAPI) GetManifestProperties(deviceTypeID string) (*model.MetadataEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/devicetypes/%s/manifest/properties", apiREST, deviceTypeID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.MetadataEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDeviceTypesInfo returns a device type device management information.
func (a *DevicesManagementAPI) GetDeviceTypesInfo(deviceTypeID string) (*model.DeviceTypeInfoEnvelope, error) {
	u := fmt.Sprintf("%s/devicemgmt/devicetypes/%s", apiREST, deviceTypeID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTypeInfoEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateDeviceTypesInfo updates a device type information.
func (a *DevicesManagementAPI) UpdateDeviceTypesInfo(deviceTypeID string, deviceTypeInfo *model.DeviceTypeInfo) (*model.DeviceTypeInfoEnvelope, error) {
	if deviceTypeInfo == nil {
		return nil, fmt.Errorf("deviceTypeInfo must be provided")
	}

	u := fmt.Sprintf("%s/devicemgmt/devicetypes/%s", apiREST, deviceTypeID)

	b, err := a.Client.NewPutRequest(u, deviceTypeInfo)
	if err != nil {
		return nil, err
	}

	result := &model.DeviceTypeInfoEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
