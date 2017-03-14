package auth

import (
	"path/filepath"

	"github.com/Unknwon/com"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/rai-project/config"
)

var (
	DefaultProfilePath string
)

func init() {
	config.AfterInit(func() {
		homeDir, err := homedir.Dir()
		if err != nil {
			return
		}

		appName := config.App.Name

		// load ~/.rai_profile
		homeProfileFile := filepath.Join(homeDir, "."+appName+"_profile")
		if com.IsFile(homeProfileFile) {
			DefaultProfilePath = homeProfileFile
			return
		}

		// load ~/.rai_env
		homeEnvFile := filepath.Join(homeDir, "."+appName+"_env")
		if com.IsFile(homeEnvFile) {
			DefaultProfilePath = homeEnvFile
			return
		}

		// load ~/.rai.profile
		homeProfileFile = filepath.Join(homeDir, "."+appName+".profile")
		if com.IsFile(homeProfileFile) {
			DefaultProfilePath = homeProfileFile
			return
		}

		// load ~/.rai.env
		homeEnvFile = filepath.Join(homeDir, "."+appName+".env")
		if com.IsFile(homeEnvFile) {
			DefaultProfilePath = homeEnvFile
			return
		}

		DefaultProfilePath = filepath.Join(homeDir, "."+appName+"_profile")
	})
}
