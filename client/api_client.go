package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// APIClient model.
type APIClient struct {
	AccessToken string
}

// NewGetRequest returns new GET http request.
func (c *APIClient) NewGetRequest(url string, values url.Values) (*[]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// auth
	c.addAuthHeader(req)

	// query
	if values != nil {
		req.URL.RawQuery = values.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &b, nil
}

// NewPostRequest returns new POST http request.
func (c *APIClient) NewPostRequest(url string, body interface{}) (*[]byte, error) {
	var payload *strings.Reader
	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		payload = strings.NewReader(string(j))
	} else {
		payload = nil
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	// auth
	c.addAuthHeader(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &b, nil
}

// NewPutRequest returns new PUT http request.
func (c *APIClient) NewPutRequest(url string, body interface{}) (*[]byte, error) {
	if body == nil {
		body = struct{}{}
	}

	j, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(j))

	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		return nil, err
	}

	// auth
	c.addAuthHeader(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &b, nil
}

// NewDeleteRequest return new DELETE http request.
func (c *APIClient) NewDeleteRequest(url string) (*[]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	// auth
	c.addAuthHeader(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &b, nil
}

func (c *APIClient) addAuthHeader(r *http.Request) {
	if c.AccessToken != "" {
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	} else {
		panic(fmt.Errorf("access token must be provided"))
	}
}

// func (c *APIClient) checkAccessToken() {
// 	if c.AccessToken == "" {
// 		panic(fmt.Errorf("access token not set"))
// 	}
// }
