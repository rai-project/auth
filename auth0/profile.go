package auth0

import (
	"strings"

	passlib "gopkg.in/hlandau/passlib.v1"

	"encoding/base64"

	"github.com/pkg/errors"
	"github.com/rai-project/auth"
	"github.com/rai-project/auth/auth0/api"
	"github.com/rai-project/config"
	"github.com/rai-project/utils"
)

// Profile ...
type Profile struct {
	api *api.Api
	auth.ProfileBase
}

// NewProfile ...
func NewProfile(opts ...auth.ProfileOption) (auth.Profile, error) {
	p, err := auth.NewProfileBase(opts...)
	if err != nil {
		return nil, err
	}
	auth0API := api.New()
	return &Profile{
		api:         auth0API,
		ProfileBase: *p,
	}, nil
}

// Create ...
func (p *Profile) Create() error {
	if p.Username == "" {
		return errors.New("username is not set")
	}
	if p.Password == "" {
		s, err := utils.EncryptString(config.App.Secret, p.Username)
		if err != nil {
			return err
		}
		p.Password = s
	}
	user, err := p.api.CreateUser(api.CreateUserRequestData{
		Username:  p.Username,
		Password:  p.Password,
		Email:     p.Email,
		GivenName: strings.TrimSpace(p.Firstname + " " + p.Lastname),
		UserMetadata: map[string]interface{}{
			"username":  p.Username,
			"firstname": p.Firstname,
			"lastname":  p.Lastname,
			"email":     p.Email,
		},
		AppMetadata: map[string]interface{}{
			"app_name": config.App.Name,
			"version":  config.App.Version,
		},
		EmailVerified: true,
		VerifyEmail:   false,
	})
	if err != nil {
		return err
	}
	p.AccessKey = user.UserID
	p.SecretKey = base64.StdEncoding.EncodeToString([]byte(p.makeSecretKey()))
	return nil
}

func (p *Profile) makeSecretKey() string {
	s, err := passlib.Hash(p.Password + ":::" + p.Username)
	if err != nil {
		log.WithError(err).
			WithField("username", p.Username).
			Error("unable to create secret key")
		return ""
	}
	return s
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

	var ep Profile
	ep.Username = p.Username
	ep.Password = p.Password
	expectedSecretKey := ep.makeSecretKey()

	if _, err := passlib.Verify(p.Password+":::"+p.Username, expectedSecretKey); err != nil {
		return false, errors.Wrap(err, "secret key did not match expected")
	}
	res, err := p.api.GetUser(api.GetUserRequestData{UserID: p.AccessKey})
	if err != nil {
		return false, errors.Wrapf(err, "unable to perform auth0 query for %v", p.AccessKey)
	}
	if res.UserID != p.AccessKey {
		return false, errors.Errorf("userids did not match. expected %v but got %v", p.AccessKey, res.UserID)
	}
	return true, nil
}

// GetByEmail ...
func (p *Profile) GetByEmail() error {
	if p.Email == "" {
		return errors.New("email is not set")
	}
	if p.Password == "" {
		s, err := utils.EncryptString(config.App.Secret, p.Username)
		if err != nil {
			return err
		}
		p.Password = s
	}

	users, err := p.api.GetUsersByEmail(p.Email)
	if err != nil {
		return err
	}
	if len(users) > 1 {
		return errors.New("More than one user with that address")
	}

	user := users[0]
	p.Username = user.Username
	p.AccessKey = user.UserID
	p.SecretKey = base64.StdEncoding.EncodeToString([]byte(p.makeSecretKey()))
	return nil
}
