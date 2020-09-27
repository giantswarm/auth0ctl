package login

import (
	"fmt"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"

	"github.com/giantswarm/auth0ctl/internal/env"
)

const (
	flagClientID     = "client-id"
	flagClientSecret = "client-secret"
	flagTenant       = "tenant"
)

type flag struct {
	ClientID     string
	ClientSecret string
	Tenant       string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.ClientID, flagClientID, "", fmt.Sprintf(`Application ID for management access. Defaults to %s environment variable.`, env.Auth0ClientID))
	cmd.Flags().StringVar(&f.ClientSecret, flagClientSecret, "", fmt.Sprintf(`Application secret for management access. Defaults to %s environment variable.`, env.Auth0ClientSecret))
	cmd.Flags().StringVar(&f.Tenant, flagTenant, "giantswarm", "Target Auth0 tenant.")
}

func (f *flag) Validate() error {
	if f.ClientID == "" {
		f.ClientID = os.Getenv(env.Auth0ClientID)
	}
	if f.ClientID == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagClientID)
	}
	if f.ClientSecret == "" {
		f.ClientSecret = os.Getenv(env.Auth0ClientSecret)
	}
	if f.ClientSecret == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagClientSecret)

	}
	if f.Tenant == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagTenant)
	}

	return nil
}
