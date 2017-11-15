package api

import (
	"encoding/json"
	"fmt"

	"github.com/projekt86/goartik-cloud/client"
	"github.com/projekt86/goartik-cloud/model"
)

// TokensAPI model.
type TokensAPI struct {
	Client *client.APIClient
}

// NewTokensAPI returns new tokens API with default API client.
func NewTokensAPI() *TokensAPI {
	return &TokensAPI{
		Client: client.DefaultAPIClient,
	}
}

// NewTokensAPIWithClient returns new tokens API with custom API client.
func NewTokensAPIWithClient(client *client.APIClient) *TokensAPI {
	return &TokensAPI{
		Client: client,
	}
}

// TokenInfo returns token info.
func (a *TokensAPI) TokenInfo() (*model.TokenInfoEnvelope, error) {
	u := fmt.Sprintf("%s/accounts/tokenInfo", apiREST)

	b, err := a.Client.NewGetRequest(u, nil)
	if err != nil {
		return nil, err
	}

	result := &model.TokenInfoEnvelope{}
	err = json.Unmarshal(*b, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
