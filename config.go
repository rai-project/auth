package auth

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type authConfig struct {
	Provider string        `json:"provider" config:"auth.provider"`
	done     chan struct{} `json:"-" config:"-"`
}

var (
	Config = &authConfig{
		done: make(chan struct{}),
	}
)

func (authConfig) ConfigName() string {
	return "Auth"
}

func (a *authConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *authConfig) Read() {
	defer close(a.done)
	vipertags.Fill(Config)
}

func (c authConfig) String() string {
	return pp.Sprintln(c)
}

func (c authConfig) Wait() {
	<-c.done
}

func (c authConfig) Debug() {
	log.Debug("Auth Config = ", c)
}

func init() {
	config.Register(Config)
}
