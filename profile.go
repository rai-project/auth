package auth

import "github.com/rai-project/model"

// Provider ...
type Provider string

// Auth0Provider ...
const (
	Auth0Provider    Provider = "auth0"
	DatabaseProvider Provider = "database"
	SecretProvider   Provider = "secret"
)

// Profile ...
type Profile interface {
	Info() ProfileBase
	Create() error
	Delete() error
	Verify() (bool, error)
	Options() ProfileOptions
	GetRole() (model.Role, error)
	FindByEmail() error
}
