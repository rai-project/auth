package auth

type Provider string

const (
	Auth0Provider  Provider = "auth0"
	SecretProvider Provider = "secret"
)

type Profile interface {
	Info() ProfileBase
	Create() error
	Verify() (bool, error)
	Options() ProfileOptions
}
