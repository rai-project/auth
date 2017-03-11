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

type ProfileOptions struct {
	Firstname   string          `json:"firstname" toml:"firstname,omitempty"`
	Lastname    string          `json:"lastname" toml:"lastname,omitempty"`
	Username    string          `json:"username" toml:"username"`
	Email       string          `json:"email" toml:"email"`
	AccessKey   string          `json:"access_key" toml:"access_key"`
	SecretKey   string          `json:"secret_key" toml:"secret_key"`
	Password    string          `json:"password" toml:"-"`
	ProfilePath string          `json:"-" toml:"-"`
	AppSecret   string          `json:"-" toml:"-"`
	Context     context.Context `json:"-" toml:"-"`
}

type ProfileOption func(*ProfileOptions)

func ProfilePath(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.ProfilePath = s
	}
}

func Username(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Username = s
	}
}

func Firstname(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Username = s
	}
}

func Lastname(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Username = s
	}
}

func Password(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Password = s
	}
}

func Email(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Email = s
	}
}

func AccessKey(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.AccessKey = s
	}
}

func SecretKey(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.SecretKey = s
	}
}

func AppSecret(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.AppSecret = s
	}
}

func NewProfileBase(iopts ...ProfileOption) (*ProfileBase, error) {

	opts := ProfileOptions{
		AppSecret:   config.App.Secret,
		ProfilePath: DefaultProfilePath,
	}

	for _, o := range iopts {
		o(&opts)
	}
	if opts.Password == "" {
		s, err := passlib.Hash(opts.AppSecret + ":::" + opts.Username)
		if err != nil {
			return nil, err
		}
		opts.Password = s
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
	if profile.Password == "" {
		profile.Password = opts.Password
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
