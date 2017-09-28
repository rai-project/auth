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

// Config ...
var (
	Config = &authConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (authConfig) ConfigName() string {
	return "Auth"
}

// SetDefaults ...
func (a *authConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *authConfig) Read() {
	defer close(a.done)
	vipertags.Fill(Config)
}

// String ...
func (c authConfig) String() string {
	return pp.Sprintln(c)
}

// Wait ...
func (c authConfig) Wait() {
	<-c.done
}

// Debug ...
func (c authConfig) Debug() {
	log.Debug("Auth Config = ", c)
}

func init() {
	config.Register(Config)
}
