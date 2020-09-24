package resourceserver

import (
	"context"
	"fmt"
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

func (r *runner) run(_ context.Context, _ *cobra.Command, _ []string) error {
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

	resourceServerExists, err := a0.ResourceServerExists(r.flag.Identifier)
	if err != nil {
		return microerror.Mask(err)
	}

	if resourceServerExists {
		fmt.Printf("Resource server with %#q identifier already exists.\n", r.flag.Identifier)
		return nil
	}

	resourceServer := &auth0.ResouceServer{
		Name:               r.flag.Name,
		Identifier:         r.flag.Identifier,
		AllowOfflineAccess: r.flag.AllowOfflineAccess,
		TokenLifetime:      r.flag.TokenLifetime,
		TokenLifetimeWeb:   r.flag.TokenLifetimeWeb,
		SigningAlgorithm:   r.flag.SigningAlgorithm,
	}

	_, err = a0.CreateResourceServer(resourceServer)
	if err != nil {
		return microerror.Mask(err)
	}

	fmt.Printf("Resource server with identifier %#q has been created.\n", r.flag.Identifier)

	return nil
}
