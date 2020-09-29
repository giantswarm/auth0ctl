package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/giantswarm/microerror"
)

func (a0 *Auth0) GetClient(clientID string) (*Client, error) {
	endpoint := fmt.Sprintf("%sclients/%s", a0.audience, clientID)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
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

	if resp.StatusCode != http.StatusOK {
		return nil, microerror.Maskf(executionFailedError, string(body))
	}

	var client *Client
	err = json.Unmarshal(body, &client)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return client, nil
}

func (a0 *Auth0) UpdateClient(clientID string, client *Client) error {
	endpoint := fmt.Sprintf("%sclients/%s", a0.audience, clientID)

	data, err := json.Marshal(client)
	if err != nil {
		return microerror.Mask(err)
	}

	req, err := http.NewRequest(http.MethodPatch, endpoint, bytes.NewReader(data))
	if err != nil {
		return microerror.Mask(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a0.accessToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := a0.httpClient.Do(req)
	if err != nil {
		return microerror.Mask(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return microerror.Mask(err)
	}

	if resp.StatusCode != http.StatusOK {
		return microerror.Maskf(executionFailedError, string(body))
	}

	return nil
}
