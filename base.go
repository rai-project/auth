package auth

import (
	"context"
	"io/ioutil"

	passlib "gopkg.in/hlandau/passlib.v1"
	yaml "gopkg.in/yaml.v2"

	"github.com/Unknwon/com"
	"github.com/pkg/errors"
	"github.com/rai-project/config"
	"github.com/rai-project/model"
)

// ProfileBase ...
type ProfileBase struct {
	ProfileOptions `json:"-" toml:"profile" yaml:"profile"`
}

// NewProfileBase ...
func NewProfileBase(iopts ...ProfileOption) (*ProfileBase, error) {

	opts := ProfileOptions{
		AppSecret:   config.App.Secret,
		ProfilePath: DefaultProfilePath,
		Context:     context.Background(),
	}

	if !com.IsFile(opts.ProfilePath) {
		err := ioutil.WriteFile(opts.ProfilePath, []byte("profile:"), 0644)
		if err != nil {
			return nil, errors.Errorf("unable to locate the profile file %v and then use it. make sure you have appropriate permissions to write to the file and/or folder", opts.ProfilePath)
		}
	}
	buf, err := ioutil.ReadFile(opts.ProfilePath)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read profile")
	}
	profile := &ProfileBase{
		ProfileOptions: opts,
	}
	if err := yaml.Unmarshal(buf, profile); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal yaml profile file")
	}

	// if .rai_profile doesn't have a team field
	if profile.Team == nil {
		profile.Team = &model.Team{}
	}

	for _, o := range iopts {
		o(&profile.ProfileOptions)
	}
	if profile.ProfileOptions.Password == "" {
		s, err := passlib.Hash(opts.AppSecret + ":::" + opts.Username)
		if err != nil {
			return nil, err
		}
		profile.ProfileOptions.Password = s
	}

	// if profile.Username == "" {
	// 	return nil, errors.New("username has not been set in auth profile")
	// }
	// if profile.Email == "" {
	// 	return nil, errors.New("email has not been set in auth profile")
	// }
	return profile, nil
}

// Info ...
func (p *ProfileBase) Info() ProfileBase {
	return *p
}

// String ...
func (p ProfileBase) String() string {
	buf, err := yaml.Marshal(p)
	if err != nil {
		log.WithError(err).Error("unable to marshal profile")
		return ""
	}
	return string(buf)
}

// Options ...
func (p *ProfileBase) Options() ProfileOptions {
	return p.ProfileOptions
}
