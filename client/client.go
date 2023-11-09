package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.wpengineapi.com/v1"
)

type apiClient struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *apiClient {
	return &apiClient{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c *apiClient) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		// Here we could map API specific error messages to more user-friendly ones or handle them accordingly
		return nil, fmt.Errorf("error, status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (c *apiClient) CreateAccountUser(accountID string, userData map[string]interface{}) (map[string]interface{}, error) {
	userEndpoint := fmt.Sprintf("%s/accounts/%s/account_users", baseURL, accountID)

	userDataBytes, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", userEndpoint, bytes.NewBuffer(userDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var user map[string]interface{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
