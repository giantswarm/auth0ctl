package client

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/microerror"
)

const (
	flagID              = "id"
	flagTenant          = "tenant"
	flagAddCallback     = "add-callback"
	flagAddLogoutURL    = "add-logout-url"
	flagAddWebOrigin    = "add-web-origin"
	flagRemoveLogoutURL = "remove-logout-url"
	flagRemoveCallback  = "remove-callback"
	flagRemoveWebOrigin = "remove-web-origin"
)

type flag struct {
	ID              string
	Tenant          string
	AddCallback     []string
	AddLogoutURL    []string
	AddWebOrigin    []string
	RemoveCallback  []string
	RemoveLogoutURL []string
	RemoveWebOrigin []string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.ID, flagID, "", "ID of the client to update.")
	cmd.Flags().StringVar(&f.Tenant, flagTenant, "giantswarm", `Auth0 tenant.`)
	cmd.Flags().StringSliceVar(&f.AddCallback, flagAddCallback, []string{}, "Callback to be added to the client.")
	cmd.Flags().StringSliceVar(&f.AddLogoutURL, flagAddLogoutURL, []string{}, "Allowed logout url to be added to the client.")
	cmd.Flags().StringSliceVar(&f.AddWebOrigin, flagAddWebOrigin, []string{}, "Web origin to be added to the client.")
	cmd.Flags().StringSliceVar(&f.RemoveCallback, flagRemoveCallback, []string{}, "Callback to be removed to the client.")
	cmd.Flags().StringSliceVar(&f.RemoveLogoutURL, flagRemoveLogoutURL, []string{}, "Allowed logout url to be removed to the client.")
	cmd.Flags().StringSliceVar(&f.RemoveWebOrigin, flagRemoveWebOrigin, []string{}, "Web origin to be removed to the client.")
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
