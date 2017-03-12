package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
)

var (
	log *logrus.Entry
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "auth/auth0/api")
	})
}
