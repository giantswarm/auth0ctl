package resourceserver

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/microerror"
)

const (
	flagIdentifier = "identifier"
	flagTenant     = "tenant"
)

type flag struct {
	Identifier string
	Tenant     string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Identifier, flagIdentifier, "", `Resource server identifier.`)
	cmd.Flags().StringVar(&f.Tenant, flagTenant, "giantswarm", `Auth0 tenant.`)

}

func (f *flag) Validate() error {
	if f.Identifier == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagIdentifier)
	}
	if f.Tenant == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagTenant)
	}

	return nil
}
