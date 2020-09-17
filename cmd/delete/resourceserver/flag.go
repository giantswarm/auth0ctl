package resourceserver

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
	cmd.Flags().StringVar(&f.Name, flagName, "", `Name of the new release. Must follow semver format.`)
}

func (f *flag) Validate() error {
	if f.Name == "" {
		return microerror.Maskf(invalidFlagError, "--%s must not be empty", flagName)
	}

	return nil
}
