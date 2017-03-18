package secret

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
	"github.com/spf13/viper"
)

type secretConfig struct {
	Provider string        `json:"provider" config:"auth.provider"`
	Secret   string        `json:"-"`
	done     chan struct{} `json:"-" config:"-"`
}

var (
	Config = &secretConfig{
		done: make(chan struct{}),
	}
)

func (secretConfig) ConfigName() string {
	return "Secret"
}

func (a *secretConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *secretConfig) Read() {
	defer close(a.done)
	vipertags.Fill(Config)
	if viper.IsSet("auth.secret") {
		a.Secret = viper.GetString("auth.secret")
	} else if config.App.Secret != "" {
		a.Secret = config.App.Secret
	} else if viper.IsSet("app.secret") {
		a.Secret = viper.GetString("app.secret")
	}
}

func (c secretConfig) Wait() {
	<-c.done
}

func (c secretConfig) String() string {
	return pp.Sprintln(c)
}

func (c secretConfig) Debug() {
	log.Debug("Secret Config = ", c)
}

func init() {
	config.Register(Config)
}
