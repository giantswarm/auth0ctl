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
