package api

import (
	"encoding/json"
	"fmt"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// RulesAPI model.
type RulesAPI struct {
	Client *client.APIClient
}

// NewRulesAPI returns new rules API with default API client.
func NewRulesAPI() *RulesAPI {
	return &RulesAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewRulesAPIWithClient returns new rules API with custom API client.
func NewRulesAPIWithClient(client *client.APIClient) *RulesAPI {
	return &RulesAPI{
		Client: client,
	}
}

// CreateRule creates a new Rule.
func (a *RulesAPI) CreateRule(ruleInfo *model.RuleInfo) (*model.RuleEnvelope, error) {
	if ruleInfo == nil {
		return nil, fmt.Errorf("info must be provided")
	}

	u := fmt.Sprintf("%s/rules", apiREST)

	b, err := a.Client.NewPostRequest(u, ruleInfo)
	if err != nil {
		return nil, err
	}

	result := &model.RuleEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteRule delete a Rule
func (a *RulesAPI) DeleteRule(ruleID string) (*model.RuleEnvelope, error) {
	u := fmt.Sprintf("%s/rules/%s", apiREST, ruleID)

	b, err := a.Client.NewDeleteRequest(u)
	if err != nil {
		return nil, err
	}

	result := &model.RuleEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetRule returns a rule using the Rule ID.
func (a *RulesAPI) GetRule(ruleID string) (*model.RuleEnvelope, error) {
	u := fmt.Sprintf("%s/rules/%s", apiREST, ruleID)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.RuleEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateRule updates an existing Rule.
func (a *RulesAPI) UpdateRule(ruleID string, ruleInfo *model.RuleInfo) (*model.RuleEnvelope, error) {
	if ruleInfo == nil {
		return nil, fmt.Errorf("info must be provided")
	}

	u := fmt.Sprintf("%s/rules/%s", apiREST, ruleID)

	b, err := a.Client.NewPutRequest(u, ruleInfo)
	if err != nil {
		return nil, err
	}

	result := &model.RuleEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
