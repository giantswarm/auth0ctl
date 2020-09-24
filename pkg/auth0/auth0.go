package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/giantswarm/microerror"

	"github.com/giantswarm/auth0ctl/internal/key"
)

var (
	managementAudience = "https://%s.eu.auth0.com/api/v2/"
)

type Config struct {
	Tenant string
}

type Auth0 struct {
	accessToken string
	audience    string

	httpClient *http.Client
}

func New(config Config) (*Auth0, error) {
	if config.Tenant == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Tenant must not be empty", config)
	}

	filePath := filepath.Join(key.ConfigDir(), config.Tenant)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var tokenConfig *TokenConfig

	err = json.Unmarshal(data, &tokenConfig)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	expiresAt, err := time.Parse(dateTimeFormat, tokenConfig.ExpiresAt)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	now, err := time.Parse(dateTimeFormat, time.Now().Format(dateTimeFormat))
	if err != nil {
		return nil, microerror.Mask(err)
	}

	now, err := time.Parse(dateTimeFormat, time.Now().Format(dateTimeFormat))
	if err != nil {
		return nil, microerror.Mask(err)
	}

	if expiresAt.Before(now) {
		return nil, microerror.Maskf(executionFailedError, "Access token expired. Execute `auth0ctl login` to get new token.")
	}

	httpClient := &http.Client{}

	return &Auth0{
		accessToken: tokenConfig.Token,
		audience:    fmt.Sprintf(managementAudience, config.Tenant),

		httpClient: httpClient,
	}, nil
}

func (a0 *Auth0) CreateResourceServer(rr *ResouceServer) (*ResouceServer, error) {
	endpoint := fmt.Sprintf("%sresource-servers", a0.audience)

	data, err := json.Marshal(rr)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(data))
	if err != nil {
		return nil, microerror.Mask(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a0.accessToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := a0.httpClient.Do(req)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	if err != nil {
		return nil, microerror.Mask(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	if resp.StatusCode == http.StatusConflict {
		return nil, microerror.Mask(resourceExistsError)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, microerror.Maskf(executionFailedError, string(body))
	}

	var newResourceServer *ResouceServer
	err = json.Unmarshal(body, &newResourceServer)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return newResourceServer, nil
}

func (a0 *Auth0) DeleteResourceServer(identifier string) error {
	endpoint := fmt.Sprintf("%sresource-servers/%s", a0.audience, url.QueryEscape(identifier))

	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return microerror.Mask(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a0.accessToken))

	resp, err := a0.httpClient.Do(req)
	if err != nil {
		return microerror.Mask(err)
	}

	if resp.StatusCode != 204 {
		return microerror.Maskf(executionFailedError, "expected status code '204', got '%d'", resp.StatusCode)
	}

	return nil

}
