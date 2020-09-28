package client

import (
	"github.com/spf13/cobra"

	"github.com/giantswarm/microerror"
)

const (
	flagName = "name"
)

type flag struct {
	Name string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Name, flagName, "", "New application name.")
}

func (f *flag) Validate() error {
	if f.Name == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagName)
	}

	return nil
}
