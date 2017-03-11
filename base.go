package auth

import (
	"context"
	"io/ioutil"

	passlib "gopkg.in/hlandau/passlib.v1"

	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/Unknwon/com"
	"github.com/pkg/errors"
	"github.com/rai-project/config"
)

type ProfileBase struct {
	ProfileOptions `json:"-" toml:"profile"`
}

func NewProfileBase(iopts ...ProfileOption) (*ProfileBase, error) {

	opts := ProfileOptions{
		AppSecret:   config.App.Secret,
		ProfilePath: DefaultProfilePath,
		Context:     context.Background(),
	}

	if !com.IsFile(opts.ProfilePath) {
		return nil, errors.Errorf("unable to locate %v. not such file or directory", opts.ProfilePath)
	}
	buf, err := ioutil.ReadFile(opts.ProfilePath)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read profile")
	}
	profile := &ProfileBase{
		ProfileOptions: opts,
	}
	_, err = toml.Decode(string(buf), profile)
	if err != nil {
		return nil, err
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

	if profile.Username == "" {
		return nil, errors.New("username has not been set in auth profile")
	}
	if profile.Email == "" {
		return nil, errors.New("email has not been set in auth profile")
	}
	return profile, nil
}

func (p *ProfileBase) Info() ProfileBase {
	return *p
}

func (p ProfileBase) String() string {
	buf := new(bytes.Buffer)
	enc := toml.NewEncoder(buf)
	enc.Encode(p)
	return buf.String()
}

func (p *ProfileBase) Options() ProfileOptions {
	return p.ProfileOptions
}
