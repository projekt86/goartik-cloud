package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// TagsAPI model.
type TagsAPI struct {
	Client *client.APIClient
}

// NewTagsAPI returns new tags API with default API client.
func NewTagsAPI() *TagsAPI {
	return &TagsAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewTagsAPIWithClient returns new tags API with custom API client.
func NewTagsAPIWithClient(client *client.APIClient) *TagsAPI {
	return &TagsAPI{
		Client: client,
	}
}

// GetAllCategories returns all tags marked as categories.
func (a *TagsAPI) GetAllCategories() (*model.TagsEnvelope, error) {
	u := fmt.Sprintf("%s/tags/categories", apiREST)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.TagsEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTagSuggestions returns tag suggestions for applications,
// device types that have been most used with a group of tags.
func (a *TagsAPI) GetTagSuggestions(name string, tags []string, entityType string, count int) (*model.TagsEnvelope, error) {
	u := fmt.Sprintf("%s/tags/suggestions", apiREST)

	// parse params
	v := url.Values{}
	if count > 0 {
		v.Add("count", parseIntParam(count, 1, 10))
	}
	v.Add("name", name)
	v.Add("entityType", entityType)
	v.Add("tags", strings.Join(tags, ","))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.TagsEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetTagsByCategories returns all tags related to the list of categories.
func (a *TagsAPI) GetTagsByCategories(categories []string) (*model.TagsEnvelope, error) {
	u := fmt.Sprintf("%s/tags/suggestions", apiREST)

	// parse params
	v := url.Values{}
	v.Add("categories", strings.Join(categories, ","))

	b, err := a.Client.NewGetRequest(u, v)
	if err != nil {
		return nil, err
	}

	result := &model.TagsEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
