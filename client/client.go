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

type ApiClient struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *ApiClient {
	return &ApiClient{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c *ApiClient) doRequest(req *http.Request) ([]byte, error) {
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

// #############################################################################
// account
// #############################################################################

func (c *ApiClient) GetAccount(accountID string) (map[string]interface{}, error) {
	accountEndpoint := fmt.Sprintf("%s/accounts/%s", baseURL, accountID)

	req, err := http.NewRequest("GET", accountEndpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var account map[string]interface{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (c *ApiClient) CreateAccount(accountData map[string]interface{}) (map[string]interface{}, error) {
	accountEndpoint := fmt.Sprintf("%s/accounts", baseURL)

	accountDataBytes, err := json.Marshal(accountData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", accountEndpoint, bytes.NewBuffer(accountDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var account map[string]interface{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (c *ApiClient) UpdateAccount(accountID string, accountData map[string]interface{}) (map[string]interface{}, error) {
	accountEndpoint := fmt.Sprintf("%s/accounts/%s", baseURL, accountID)

	accountDataBytes, err := json.Marshal(accountData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", accountEndpoint, bytes.NewBuffer(accountDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var account map[string]interface{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (c *ApiClient) DeleteAccount(accountID string) error {
	accountEndpoint := fmt.Sprintf("%s/accounts/%s", baseURL, accountID)

	req, err := http.NewRequest("DELETE", accountEndpoint, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

// end account

// #############################################################################
// account_user
// #############################################################################

func (c *ApiClient) CreateAccountUser(accountID string, userData map[string]interface{}) (map[string]interface{}, error) {
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

func (c *ApiClient) GetAccountUser(userID string) (map[string]interface{}, error) {
	userEndpoint := fmt.Sprintf("%s/users/%s", baseURL, userID)

	req, err := http.NewRequest("GET", userEndpoint, nil)
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

func (c *ApiClient) UpdateAccountUser(userID string, userData map[string]interface{}) (map[string]interface{}, error) {
	userEndpoint := fmt.Sprintf("%s/users/%s", baseURL, userID)

	userDataBytes, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", userEndpoint, bytes.NewBuffer(userDataBytes))
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

func (c *ApiClient) DeleteAccountUser(userID string) error {
	userEndpoint := fmt.Sprintf("%s/users/%s", baseURL, userID)

	req, err := http.NewRequest("DELETE", userEndpoint, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

// end account_user

// #############################################################################
// cdn
// #############################################################################

func (c *ApiClient) GetCDN(cdnID string) (map[string]interface{}, error) {
	cdnEndpoint := fmt.Sprintf("%s/cdns/%s", baseURL, cdnID)

	req, err := http.NewRequest("GET", cdnEndpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var cdn map[string]interface{}
	err = json.Unmarshal(body, &cdn)
	if err != nil {
		return nil, err
	}

	return cdn, nil
}

func (c *ApiClient) CreateCDN(cdnData map[string]interface{}) (map[string]interface{}, error) {
	cdnEndpoint := fmt.Sprintf("%s/cdns", baseURL)

	cdnDataBytes, err := json.Marshal(cdnData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", cdnEndpoint, bytes.NewBuffer(cdnDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var cdn map[string]interface{}
	err = json.Unmarshal(body, &cdn)
	if err != nil {
		return nil, err
	}

	return cdn, nil
}

func (c *ApiClient) UpdateCDN(cdnID string, cdnData map[string]interface{}) (map[string]interface{}, error) {
	cdnEndpoint := fmt.Sprintf("%s/cdns/%s", baseURL, cdnID)

	cdnDataBytes, err := json.Marshal(cdnData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", cdnEndpoint, bytes.NewBuffer(cdnDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var cdn map[string]interface{}
	err = json.Unmarshal(body, &cdn)
	if err != nil {
		return nil, err
	}

	return cdn, nil
}

func (c *ApiClient) DeleteCDN(cdnID string) error {
	cdnEndpoint := fmt.Sprintf("%s/cdns/%s", baseURL, cdnID)

	req, err := http.NewRequest("DELETE", cdnEndpoint, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

// end cdn

// #############################################################################
// domain
// #############################################################################
// GetDomain retrieves details of a specific domain.
func (c *ApiClient) GetDomain(domainID string) (map[string]interface{}, error) {
	domainEndpoint := fmt.Sprintf("%s/domains/%s", baseURL, domainID)

	req, err := http.NewRequest("GET", domainEndpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var domain map[string]interface{}
	err = json.Unmarshal(body, &domain)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

// CreateDomain sets up a new domain configuration.
func (c *ApiClient) CreateDomain(domainData map[string]interface{}) (map[string]interface{}, error) {
	domainEndpoint := fmt.Sprintf("%s/domains", baseURL)

	domainDataBytes, err := json.Marshal(domainData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", domainEndpoint, bytes.NewBuffer(domainDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var domain map[string]interface{}
	err = json.Unmarshal(body, &domain)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

// UpdateDomain modifies a specific domain configuration.
func (c *ApiClient) UpdateDomain(domainID string, domainData map[string]interface{}) (map[string]interface{}, error) {
	domainEndpoint := fmt.Sprintf("%s/domains/%s", baseURL, domainID)

	domainDataBytes, err := json.Marshal(domainData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", domainEndpoint, bytes.NewBuffer(domainDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var domain map[string]interface{}
	err = json.Unmarshal(body, &domain)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

// DeleteDomain removes a specific domain configuration.
func (c *ApiClient) DeleteDomain(domainID string) error {
	domainEndpoint := fmt.Sprintf("%s/domains/%s", baseURL, domainID)

	req, err := http.NewRequest("DELETE", domainEndpoint, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

// end domain

// #############################################################################
// install
// #############################################################################

func (c *ApiClient) GetInstall(installID string) (map[string]interface{}, error) {
	installEndpoint := fmt.Sprintf("%s/installs/%s", baseURL, installID)

	req, err := http.NewRequest("GET", installEndpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var install map[string]interface{}
	err = json.Unmarshal(body, &install)
	if err != nil {
		return nil, err
	}

	return install, nil
}

func (c *ApiClient) CreateInstall(installData map[string]interface{}) (map[string]interface{}, error) {
	installEndpoint := fmt.Sprintf("%s/installs", baseURL)

	installDataBytes, err := json.Marshal(installData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", installEndpoint, bytes.NewBuffer(installDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var install map[string]interface{}
	err = json.Unmarshal(body, &install)
	if err != nil {
		return nil, err
	}

	return install, nil
}

func (c *ApiClient) UpdateInstall(installID string, installData map[string]interface{}) (map[string]interface{}, error) {
	installEndpoint := fmt.Sprintf("%s/installs/%s", baseURL, installID)

	installDataBytes, err := json.Marshal(installData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", installEndpoint, bytes.NewBuffer(installDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var install map[string]interface{}
	err = json.Unmarshal(body, &install)
	if err != nil {
		return nil, err
	}

	return install, nil
}

func (c *ApiClient) DeleteInstall(installID string) error {
	installEndpoint := fmt.Sprintf("%s/installs/%s", baseURL, installID)

	req, err := http.NewRequest("DELETE", installEndpoint, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

// end install

// #############################################################################
// site
// #############################################################################

func (c *ApiClient) GetSite(siteID string) (map[string]interface{}, error) {
	siteEndpoint := fmt.Sprintf("%s/sites/%s", baseURL, siteID)

	req, err := http.NewRequest("GET", siteEndpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var site map[string]interface{}
	err = json.Unmarshal(body, &site)
	if err != nil {
		return nil, err
	}

	return site, nil
}

func (c *ApiClient) CreateSite(siteData map[string]interface{}) (map[string]interface{}, error) {
	siteEndpoint := fmt.Sprintf("%s/sites", baseURL)

	siteDataBytes, err := json.Marshal(siteData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", siteEndpoint, bytes.NewBuffer(siteDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var site map[string]interface{}
	err = json.Unmarshal(body, &site)
	if err != nil {
		return nil, err
	}

	return site, nil
}

func (c *ApiClient) UpdateSite(siteID string, siteData map[string]interface{}) (map[string]interface{}, error) {
	siteEndpoint := fmt.Sprintf("%s/sites/%s", baseURL, siteID)

	siteDataBytes, err := json.Marshal(siteData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", siteEndpoint, bytes.NewBuffer(siteDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var site map[string]interface{}
	err = json.Unmarshal(body, &site)
	if err != nil {
		return nil, err
	}

	return site, nil
}

func (c *ApiClient) DeleteSite(siteID string) error {
	siteEndpoint := fmt.Sprintf("%s/sites/%s", baseURL, siteID)

	req, err := http.NewRequest("DELETE", siteEndpoint, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

// end site

// #############################################################################
// ssh_key
// #############################################################################

func (c *ApiClient) GetSSHKey(sshKeyID string) (map[string]interface{}, error) {
	sshKeyEndpoint := fmt.Sprintf("%s/ssh_keys/%s", baseURL, sshKeyID)

	req, err := http.NewRequest("GET", sshKeyEndpoint, nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var sshKey map[string]interface{}
	err = json.Unmarshal(body, &sshKey)
	if err != nil {
		return nil, err
	}

	return sshKey, nil
}

func (c *ApiClient) CreateSSHKey(sshKeyData map[string]interface{}) (map[string]interface{}, error) {
	sshKeyEndpoint := fmt.Sprintf("%s/ssh_keys", baseURL)

	sshKeyDataBytes, err := json.Marshal(sshKeyData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", sshKeyEndpoint, bytes.NewBuffer(sshKeyDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var sshKey map[string]interface{}
	err = json.Unmarshal(body, &sshKey)
	if err != nil {
		return nil, err
	}

	return sshKey, nil
}

func (c *ApiClient) UpdateSSHKey(sshKeyID string, sshKeyData map[string]interface{}) (map[string]interface{}, error) {
	sshKeyEndpoint := fmt.Sprintf("%s/ssh_keys/%s", baseURL, sshKeyID)

	sshKeyDataBytes, err := json.Marshal(sshKeyData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", sshKeyEndpoint, bytes.NewBuffer(sshKeyDataBytes))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var sshKey map[string]interface{}
	err = json.Unmarshal(body, &sshKey)
	if err != nil {
		return nil, err
	}

	return sshKey, nil
}

func (c *ApiClient) DeleteSSHKey(sshKeyID string) error {
	sshKeyEndpoint := fmt.Sprintf("%s/ssh_keys/%s", baseURL, sshKeyID)

	req, err := http.NewRequest("DELETE", sshKeyEndpoint, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	return err
}

// end ssh_key
