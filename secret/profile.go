package secret

import (
	"errors"

	"github.com/rai-project/acl"
	"github.com/rai-project/auth"
)

// Profile ...
type Profile struct {
	*auth.ProfileBase
}

// NewProfile ...
func NewProfile(opts ...auth.ProfileOption) (auth.Profile, error) {
	p, err := auth.NewProfileBase(opts...)
	if err != nil {
		return nil, err
	}
	return &Profile{p}, nil
}

// Create ...
func (p *Profile) Create() error {
	if p.Username == "" {
		return errors.New("username is not set")
	}
	accessKeyHash, secretKeyHash, err := Hash(p.Username)
	if err != nil {
		return err
	}
	p.AccessKey = accessKeyHash
	p.SecretKey = secretKeyHash
	return nil
}

// Verify ...
func (p *Profile) Verify() (bool, error) {
	if p.Username == "" {
		return false, errors.New("username is not set")
	}
	if p.AccessKey == "" {
		return false, errors.New("access key is not set")
	}
	if p.SecretKey == "" {
		return false, errors.New("secret key is not set")
	}
	return Verify(p.Username, p.AccessKey, p.SecretKey)
}

func (p *Profile) GetRole() (acl.Role, error) {
	panic("GetRole not implemented")
	return acl.Role(""), nil
}

// FindByEmail ...
func (p *Profile) FindByEmail() error {
	panic("FindByEmail unimplemented")
	return nil
}

// Delete ...
func (p *Profile) Delete() error {
	panic("Delete unimplemented")
	return nil
}
