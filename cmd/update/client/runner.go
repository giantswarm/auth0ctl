package client

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

	newCallbacks := r.flag.AddCallback
	{
		for _, callback := range client.Callbacks {
			if !stringInSlice(callback, r.flag.RemoveCallback) {
				newCallbacks = append(newCallbacks, callback)
			}
		}
	}
	client.Callbacks = newCallbacks

	newAllowedLogoutURLs := r.flag.AddAllowedLogoutURL
	{
		for _, allowedLogoutURL := range client.AllowedLogoutURLs {
			if !stringInSlice(allowedLogoutURL, r.flag.RemoveAllowedLogoutURL) {
				newAllowedLogoutURLs = append(newAllowedLogoutURLs, allowedLogoutURL)
			}
		}
	}
	client.AllowedLogoutURLs = newAllowedLogoutURLs

	newWebOrigins := r.flag.AddWebOrigin
	{
		for _, webOrigin := range client.WebOrigins {
			if !stringInSlice(webOrigin, r.flag.RemoveWebOrigin) {
				newWebOrigins = append(newWebOrigins, webOrigin)
			}
		}
	}
	client.WebOrigins = newWebOrigins

	err = a0.UpdateClient(r.flag.ID, client)
	if err != nil {
		return microerror.Mask(err)
	}

	fmt.Printf("Auth0 client with ID %#q in tenant %#q  was successfully updated.\n", r.flag.ID, r.flag.Tenant)

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
