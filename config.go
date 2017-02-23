package auth

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type authConfig struct {
	Secret string `json:"secret" config:"app.secret" default:"Hammurabi"`
}

var (
	Config = &authConfig{}
)

func (authConfig) ConfigName() string {
	return "Auth"
}

func (authConfig) SetDefaults() {
}

func (a *authConfig) Read() {
	vipertags.Fill(a)
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
