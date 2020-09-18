package resourceserver

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/microerror"
)

const (
	flagName               = "name"
	flagIdentifier         = "identifier"
	flagTenant             = "tenant"
	flagAllowOfflineAccess = "allow-offline-access"
	flagTokenLifetime      = "token-lifetime"
	flagTokenLifetimeWeb   = "token-lifetime-web"
	flagSigningAlgorithm   = "signing-algorithm"
)

type flag struct {
	Name               string
	Identifier         string
	Tenant             string
	AllowOfflineAccess bool
	TokenLifetime      int
	TokenLifetimeWeb   int
	SigningAlgorithm   string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Name, flagName, "", `Name of a new resource server.`)
	cmd.Flags().StringVar(&f.Identifier, flagIdentifier, "", `Resource server identifier.`)
	cmd.Flags().StringVar(&f.Tenant, flagTenant, "giantswarm", `Auth0 tenant.`)
	cmd.Flags().BoolVar(&f.AllowOfflineAccess, flagAllowOfflineAccess, true, `Allow offline access.`)
	cmd.Flags().IntVar(&f.TokenLifetime, flagTokenLifetime, 300, `Token lifetime.`)
	cmd.Flags().IntVar(&f.TokenLifetimeWeb, flagTokenLifetimeWeb, 300, `Token lifetime for web`)
	cmd.Flags().StringVar(&f.SigningAlgorithm, flagSigningAlgorithm, "RS256", `Signing algorithm.`)
}

func (f *flag) Validate() error {
	if f.Name == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagName)
	}
	if f.Identifier == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagIdentifier)
	}
	if f.Tenant == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagTenant)
	}
	if f.SigningAlgorithm != "HS256" && f.SigningAlgorithm != "RS256" {
		return microerror.Maskf(invalidFlagError, "--%s must be one of ['HS256', 'RS256']", flagSigningAlgorithm)
	}

	return nil
}
