package auth

import (
	"context"
	"strings"

	"github.com/rai-project/model"
)

// ProfileOptions ...
type ProfileOptions struct {
	model.User  `toml:"" yaml:",inline"`
	ProfilePath string          `json:"-" yaml:"-" toml:"-"`
	AppSecret   string          `json:"-" yaml:"-" toml:"-"`
	Context     context.Context `json:"-" yaml:"-" toml:"-"`
}

// ProfileOption ...
type ProfileOption func(*ProfileOptions)

// ProfilePath ...
func ProfilePath(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.ProfilePath = s
	}
}

// Username ...
func Username(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Username = strings.ToLower(s)
	}
}

// Firstname ...
func Firstname(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Firstname = strings.Title(s)
	}
}

// Lastname ...
func Lastname(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Lastname = strings.Title(s)
	}
}

// Password ...
func Password(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Password = s
	}
}

// Email ...
func Email(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Email = s
	}
}

// Affiliation ...
func Affiliation(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Affiliation = s
	}
}

// AccessKey ...
func AccessKey(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.AccessKey = s
	}
}

// SecretKey ...
func SecretKey(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.SecretKey = s
	}
}

// AppSecret ...
func AppSecret(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.AppSecret = s
	}
}

// TeamName ...
func TeamName(s string) ProfileOption {
	return func(o *ProfileOptions) {
		if o.Team == nil {
			o.Team = &model.Team{}
		}
		o.Team.Name = s
	}
}
