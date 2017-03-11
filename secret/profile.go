package secret

import (
	"errors"

	"github.com/rai-project/auth"
)

type Profile struct {
	*auth.ProfileBase
}

func NewProfile(opts ...auth.ProfileOption) (auth.Profile, error) {
	p, err := auth.NewProfileBase(opts...)
	if err != nil {
		return nil, err
	}
	return &Profile{p}, nil
}

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