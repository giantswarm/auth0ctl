package auth0

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/auth0ctl/internal/key"
)

const (
	dateTimeFormat = "2006-01-02T15:04:05"
)

type ClientCredentials struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int16  `json:"expires_in"`
}

type TokenConfig struct {
	ExpiresAt string `json:"expires_at"`
	Token     string `json:"token"`
}

func Login(clientID, clientSecret, tenant string) error {
	err := ensureConfigDirExists()
	if err != nil {
		return microerror.Mask(err)
	}

	filePath := filepath.Join(key.ConfigDir(), tenant)

	clientCredentials, err := getAccessToken(clientID, clientSecret, tenant)
	if err != nil {
		return microerror.Mask(err)
	}
	ttl := time.Second * time.Duration(clientCredentials.ExpiresIn)

	expiresAt := time.Now().Add(ttl).Format(dateTimeFormat)

	tokenConfig := TokenConfig{
		Token:     clientCredentials.AccessToken,
		ExpiresAt: expiresAt,
	}

	data, err := json.Marshal(tokenConfig)
	if err != nil {
		return microerror.Mask(err)
	}

	err = ioutil.WriteFile(filePath, tokenConfig, 0600)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func Logout(tenant string) error {
	filePath := filepath.Join(key.ConfigDir(), tenant)

	err := os.Remove(filePath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func ensureConfigDirExists() error {
	_, err := os.Stat(key.ConfigDir())
	if os.IsExist(err) {
		return nil
	}

	err = os.MkdirAll(key.ConfigDir(), 0744)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func getAccessToken(clientID, clientSecret, tenant string) (*ClientCredentials, error) {
	httpClient := &http.Client{}

	authEndpoint := fmt.Sprintf("https://%s.eu.auth0.com/oauth/token", tenant)

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("audience", fmt.Sprintf(managementAudience, tenant))
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", authEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, microerror.Mask(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var clientCredentials *ClientCredentials

	err = json.Unmarshal(body, &clientCredentials)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return clientCredentials, nil
}

func writeTokenConfigToFileSystem(tokenConfig []byte, filePath string) error {

	err := ioutil.WriteFile(filePath, tokenConfig, 0600)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
