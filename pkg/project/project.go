package project

var (
	description = "Command line tool for Auth0."
	gitSHA      = "n/a"
	name        = "auth0ctl"
	source      = "https://github.com/giantswarm/auth0ctl"
	version     = "0.1.1-dev"
)

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
