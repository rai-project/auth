package provider

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/rai-project/auth"
	"github.com/rai-project/auth/auth0"
	"github.com/rai-project/auth/secret"
)

// New ...
func New(opts ...auth.ProfileOption) (auth.Profile, error) {
	provider := auth.Provider(strings.ToLower(auth.Config.Provider))
	switch provider {
	case auth.Auth0Provider:
		return auth0.NewProfile(opts...)
	case auth.SecretProvider:
		return secret.NewProfile(opts...)
	case auth.DatabaseProvider:

	default:
		return nil, errors.Errorf("the auth provider %v specified is not supported", provider)
	}
}
