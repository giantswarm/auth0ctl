package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/giantswarm/microerror"
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

	accessToken, err := readTokenFromFile(config.Tenant)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	httpClient := &http.Client{}

	return &Auth0{
		accessToken: accessToken,
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

	fmt.Println(string(data))

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
		return microerror.Maskf(executionFailedError, "expected status code '204', got %#d", resp.StatusCode)
	}

	return nil

}

func (a0 *Auth0) ResourceServerExists(identifier string) (bool, error) {
	endpoint := fmt.Sprintf("%sresource-servers/%s", a0.audience, url.QueryEscape(identifier))

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return false, microerror.Mask(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a0.accessToken))

	resp, err := a0.httpClient.Do(req)
	if err != nil {
		return false, microerror.Mask(err)
	}

	if resp.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}
