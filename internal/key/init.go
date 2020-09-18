package key

import (
	"os/user"

	"github.com/giantswarm/microerror"
)

func init() {
	var err error

	// Initialize osUser.
	{
		osUser, err = user.Current()
		if err != nil {
			panic(microerror.JSON(err))
		}
	}
}
