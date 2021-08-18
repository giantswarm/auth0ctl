[![CircleCI](https://circleci.com/gh/giantswarm/auth0ctl.svg?style=shield&circle-token=5f432129bee4f3b1d8a875c5c2bf8aed0cda6bea)](https://circleci.com/gh/giantswarm/auth0ctl)

# auth0ctl

Command line client for Auth0.

## Installation

```sh
make install
```

## Configuration

`auth0ctl` requires credentials with access to [Management API](https://auth0.com/docs/api/management/v2).

Use environment variables to configure cli:
- `AUTH0_CLIENT_ID`: client ID of the Auth0 application.
- `AUTH0_CLIENT_SECRET`: client secret of the Auth0 application.

Required application scopes:
  - read:clients
  - read:client_keys
  - update:client_keys
  - create:clients
  - delete:clients
  - update:clients
  - read:resource_servers
  - create:resource_servers
  - delete:resource_servers

## Usage

Please check `auth0ctl -h` for for details on all functions.
