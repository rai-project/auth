package auth0

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type auth0Config struct {
	Provider     string `json:"provider" config:"auth.provider" default:"auth0"`
	Domain       string `json:"domain" config:"auth.domain"`
	ClientID     string `json:"-" config:"auth.client_id"`
	ClientSecret string `json:"-" config:"auth.client_secret"`
	Connection   string `json:"connection" config:"auth.connection" default:"Username-Password-Authentication"`
}

var (
	Config = &auth0Config{}
)

func (auth0Config) ConfigName() string {
	return "Auth0"
}

func (auth0Config) SetDefaults() {
}

func (a *auth0Config) Read() {
	vipertags.Fill(a)
}

func (c auth0Config) String() string {
	return pp.Sprintln(c)
}

func (c auth0Config) Debug() {
	log.Debug("Auth0 Config = ", c)
}

func init() {
	config.Register(Config)
}
