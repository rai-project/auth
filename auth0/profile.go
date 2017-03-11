package auth0

import (
	"encoding/base64"

	"github.com/fatih/structs"
	"github.com/pkg/errors"
	"github.com/rai-project/auth"
	"github.com/rai-project/auth/auth0/api"
	"github.com/rai-project/config"
	"github.com/rai-project/utils"
)

type Profile struct {
	api *api.Api
	*auth.ProfileBase
}

func NewProfile(opts ...auth.ProfileOption) (auth.Profile, error) {
	p, err := auth.NewProfileBase(opts...)
	if err != nil {
		return nil, err
	}
	auth0API := api.New()
	return &Profile{
		api:         auth0API,
		ProfileBase: p,
	}, nil
}

func (p *Profile) Create() error {
	if p.Username == "" {
		return errors.New("username is not set")
	}
	user, err := p.api.CreateUser(api.CreateUserRequestData{
		Username:    p.Username,
		Email:       p.Email,
		AppMetadata: structs.Map(config.App),
	})
	if err != nil {
		return err
	}
	p.AccessKey = user.UserID
	p.SecretKey = p.makeSecretKey()
	return nil
}

func (p *Profile) makeSecretKey() string {
	s, err := utils.EncryptString(config.App.Secret, p.Username)
	if err != nil {
		log.WithError(err).
			WithField("username", p.Username).
			Error("unable to create secret key")
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(s))
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
	secretKey := p.makeSecretKey()
	ep := Profile{ProfileBase: &auth.ProfileBase{Username: p.Username}}
	expectedSecretKey := ep.makeSecretKey()
	if secretKey != expectedSecretKey {
		return false, errors.New("secret key did not match expected")
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
