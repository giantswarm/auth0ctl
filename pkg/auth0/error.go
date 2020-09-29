package auth0

import "github.com/giantswarm/microerror"

var resourceExistsError = &microerror.Error{
	Kind: "resourceExistsError",
}

// IsResourceExists asserts resourceExistsError.
func IsResourceExists(err error) bool {
	return microerror.Cause(err) == resourceExistsError
}

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var executionFailedError = &microerror.Error{
	Kind: "executionFailedError",
}

// IsInvalidFlags asserts invalidFlagsError.
func IsExecutionFailedError(err error) bool {
	return microerror.Cause(err) == executionFailedError
}
