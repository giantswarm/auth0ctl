package client

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/microerror"
)

const (
	flagID     = "id"
	flagTenant = "tenant"
)

type flag struct {
	ID     string
	Tenant string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.ID, flagID, "", "ID of the client to update.")
	cmd.Flags().StringVar(&f.Tenant, flagTenant, "giantswarm", `Auth0 tenant.`)
}

func (f *flag) Validate() error {
	if f.ID == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagID)
	}
	if f.Tenant == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagTenant)
	}

	return nil
}
