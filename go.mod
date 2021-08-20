module github.com/giantswarm/auth0ctl

go 1.16

require (
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	github.com/spf13/cobra v1.2.1
)

replace (
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
)
