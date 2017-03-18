package api

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/utils"
	"github.com/rai-project/vipertags"
)

type auth0Config struct {
	Provider     string        `json:"provider" config:"auth.provider"`
	Domain       string        `json:"domain" config:"auth.domain"`
	ClientID     string        `json:"-" config:"auth.client_id"`
	ClientSecret string        `json:"-" config:"auth.client_secret"`
	Connection   string        `json:"connection" config:"auth.connection" default:"Username-Password-Authentication"`
	done         chan struct{} `json:"-" config:"-"`
}

var (
	Config = &auth0Config{
		done: make(chan struct{}),
	}
)

func (auth0Config) ConfigName() string {
	return "Auth0"
}

func (a *auth0Config) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *auth0Config) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if utils.IsEncryptedString(a.ClientID) {
		s, err := utils.DecryptStringBase64(config.App.Secret, a.ClientID)
		if err == nil {
			a.ClientID = s
		}
	}
	if utils.IsEncryptedString(a.ClientSecret) {
		s, err := utils.DecryptStringBase64(config.App.Secret, a.ClientSecret)
		if err == nil {
			a.ClientSecret = s
		}
	}
}

func (c auth0Config) Wait() {
	<-c.done
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
