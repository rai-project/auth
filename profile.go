package auth

import "github.com/rai-project/acl"

type Provider string

// Providers
const (
	Auth0Provider    Provider = "auth0"
	DatabaseProvider Provider = "database"
	SecretProvider   Provider = "secret"
)

// Profiles are in effect users for rai
// each user has a profile that's backed
// by a profile provider
type Profile interface {
	// Create a profile using the provider specified
	Create() error
	// Delete a profile
	Delete() error
	// Information on the profile
	Info() ProfileBase
	// Verify the profile by looking
	// at the provider and checking if the
	// credentials match
	Verify() (bool, error)
	// Profile options
	Options() ProfileOptions
	// Get the role of the user
	// roles are essentially permission
	// groups
	GetRole() (acl.Role, error)
	// Find a profile by an email. The
	// profile is found by querying the
	// profile provier
	FindByEmail() error
}
