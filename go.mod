module github.com/giantswarm/auth0ctl

go 1.19

require (
	github.com/giantswarm/microerror v0.4.1
	github.com/giantswarm/micrologger v1.1.1
	github.com/spf13/cobra v1.8.0
)

require (
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace (
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.5.1
)
