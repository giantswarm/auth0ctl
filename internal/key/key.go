package key

import (
	"os/user"
	"path"
)

// osUser is initialized in init() in init.go.
var osUser *user.User

func ConfigDir() string {
	return path.Join(osUser.HomeDir, ".config/auth0ctl")
}

func HomeDir() string {
	return osUser.HomeDir
}

func Username() string {
	return osUser.Username
}
