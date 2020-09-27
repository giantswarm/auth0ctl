package client

import (
	"context"
	"io"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/spf13/cobra"

	"github.com/giantswarm/auth0ctl/pkg/auth0"
)

type runner struct {
	flag   *flag
	logger micrologger.Logger
	stdout io.Writer
	stderr io.Writer
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.flag.Validate()
	if err != nil {
		return microerror.Mask(err)
	}

	err = r.run(ctx, cmd, args)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	var err error

	var a0 *auth0.Auth0
	{
		c := auth0.Config{
			Tenant: r.flag.Tenant,
		}

		a0, err = auth0.New(c)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	client, err := a0.GetClient(r.flag.ID)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
