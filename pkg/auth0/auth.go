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

	"github.com/giantswarm/auth0ctl/internal/key"
	"github.com/giantswarm/microerror"
)

func Login(clientID, clientSecret, tenant string) error {
	err := ensureConfigDirExists()
	if err != nil {
		return microerror.Mask(err)
	}

	accessToken, err := getAccessToken(clientID, clientSecret, tenant)
	if err != nil {
		return microerror.Mask(err)
	}

	filePath := filepath.Join(key.ConfigDir(), tenant)

	err = writeTokenToFileSystem(accessToken, filePath)
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

func getAccessToken(clientID, clientSecret, tenant string) (string, error) {
	httpClient := &http.Client{}

	authEndpoint := fmt.Sprintf("https://%s.eu.auth0.com/oauth/token", tenant)

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("audience", fmt.Sprintf(managementAudience, tenant))
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", authEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", microerror.Mask(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	type auth0LoginResponse struct {
		AccessToken string `json:"access_token"`
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", microerror.Mask(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", microerror.Mask(err)
	}

	var auth0LoginData auth0LoginResponse

	err = json.Unmarshal(body, &auth0LoginData)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return auth0LoginData.AccessToken, nil
}

func writeTokenToFileSystem(accessToken, filePath string) error {
	err := ioutil.WriteFile(filePath, []byte(accessToken), 0600)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
