module github.com/giantswarm/auth0ctl

go 1.19

require (
	github.com/giantswarm/microerror v0.4.0
	github.com/giantswarm/micrologger v1.0.0
	github.com/spf13/cobra v1.6.0
)

require (
	github.com/go-kit/log v0.2.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace (
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.5.0
)
