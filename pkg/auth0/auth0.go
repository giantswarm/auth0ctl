package auth0

import (
	"fmt"
	"net/http"

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

	accessToken, err := readTokenFromFile(config.Tenant)
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

func readTokenFromFile(tenant string) (string, error) {
	return "", nil
}
