package auth0

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/giantswarm/microerror"
)

var (
	managementAudience = "https://%s.eu.auth0.com/api/v2/"
)

type Config struct {
	// credentials
	ClientID     string
	ClientSecret string

	Tenant string
}

type Auth0 struct {
	AccessToken string
	Audience    string

	httpClient *http.Client
}

func New(config Config) (*Auth0, error) {

	if config.ClientID == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ClientID must not be empty", config)
	}
	if config.ClientSecret == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.ClientSecret must not be empty", config)
	}
	if config.Tenant == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Tenant must not be empty", config)
	}

	accessToken, err := getAccessToken(config.ClientID, config.ClientSecret, config.Tenant)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	httpClient := &http.Client{}

	return &Auth0{
		AccessToken: accessToken,
		Audience:    fmt.Sprintf(managementAudience, config.Tenant),

		httpClient: httpClient,
	}, nil
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
