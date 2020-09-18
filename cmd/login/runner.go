package login

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

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	err := auth0.Login(r.flag.ClientID, r.flag.ClientSecret, r.flag.Tenant)
	if err != nil {
		return microerror.Mask(err)
	}

	fmt.Printf("Successfully logged into Auth0 tenant %#q\n", r.flag.Tenant)

	return nil
}
