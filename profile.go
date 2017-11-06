package auth

// Provider ...
type Provider string

// Auth0Provider ...
const (
	Auth0Provider  Provider = "auth0"
	SecretProvider Provider = "secret"
)

// Profile ...
type Profile interface {
	Info() ProfileBase
	Create() error
	Verify() (bool, error)
	Options() ProfileOptions
	GetByEmail() error
}
