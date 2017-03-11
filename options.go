package auth

import (
	"context"
	"strings"

	"github.com/rai-project/model"
)

type ProfileOptions struct {
	model.User
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
		o.Username = strings.ToLower(s)
	}
}

func Firstname(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Firstname = strings.Title(s)
	}
}

func Lastname(s string) ProfileOption {
	return func(o *ProfileOptions) {
		o.Lastname = strings.Title(s)
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
