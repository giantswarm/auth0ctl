package logout

import (
	"github.com/giantswarm/microerror"
	"github.com/spf13/cobra"
)

const (
	flagTenant = "tenant"
)

type flag struct {
	Tenant string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Tenant, flagTenant, "giantswarm", "Target Auth0 tenant.")
}

func (f *flag) Validate() error {
	if f.Tenant == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagTenant)
	}

	return nil
}
