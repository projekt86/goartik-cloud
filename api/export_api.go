package api

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// ExportAPI model.
type ExportAPI struct {
	Client *client.APIClient
}

// NewExportAPI returns new export API with default API client.
func NewExportAPI() *ExportAPI {
	return &ExportAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewExportAPIWithClient return export API with custom API client.
func NewExportAPIWithClient(client *client.APIClient) *ExportAPI {
	return &ExportAPI{
		Client: client,
	}
}

// CreateExportRequest exports normalized messages.
func (a *ExportAPI) CreateExportRequest(requestInfo *model.ExportRequestInfo) (*model.ExportRequestEnvelope, error) {
	if requestInfo == nil {
		return nil, fmt.Errorf("requestInfo must be provided")
	}

	u := fmt.Sprintf("%s/messages/export", apiREST)

	b, err := a.Client.NewPostRequest(u, requestInfo)
	if err != nil {
		return nil, err
	}

	result := &model.ExportRequestEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetExportHistory returns the history of export requests.
func (a *ExportAPI) GetExportHistory(trialID string, count int, offset int) (*model.ExportRequestEnvelope, error) {
	u := fmt.Sprintf("%s/messages/export/history", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 100))
	}
	if offset > 0 {
		v.Add("offset", parseIntParam(offset, 1, math.MaxInt32))
	}
	v.Add("trialId", trialID)

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.ExportRequestEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CheckExportStatus check status of the export query.
func (a *ExportAPI) CheckExportStatus(exportID string) (*model.ExportStatusEnvelope, error) {
	u := fmt.Sprintf("%s/messages/export/%s/status", apiREST, exportID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.ExportStatusEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// // CheckExportResult : check export result
// func (a *ExportAPI) CheckExportResult(exportID string) (*model.ExportResultEnvelope, error) {
// 	u := fmt.Sprintf("%s/messages/export/%s/result", apiREST, exportID)

// 	b, err := a.Client.NewGetRequest(u, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	result := &model.ExportResultEnvelope{}
// 	err = json.Unmarshal(*b, result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil
// }
