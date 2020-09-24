package auth0

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/giantswarm/auth0ctl/internal/key"
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
