package auth

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type authConfig struct {
	Provider string `json:"provider" config:"auth.provider"`
}

var (
	Config = &authConfig{}
)

func (authConfig) ConfigName() string {
	return "Auth"
}

func (a *authConfig) SetDefaults() {
}

func (a *authConfig) Read() {
	vipertags.Fill(Config)
}

func (c authConfig) String() string {
	return pp.Sprintln(c)
}

func (c authConfig) Debug() {
	log.Debug("Auth Config = ", c)
}

func init() {
	config.Register(Config)
}
