package secret

import (
	"github.com/pkg/errors"
	passlib "github.com/rai-project/passlib"
)

// Hash ...
func Hash(username string) (accessKeyHash string, secretKeyHash string, err error) {
	accessKeyHash, err = passlib.Hash(Config.Secret + ":::" + username)
	if err != nil {
		return
	}
	secretKeyHash, err = passlib.Hash(Config.Secret + ":::" + accessKeyHash)
	return
}

// Verify ...
func Verify(username, accessKeyHash, secretKeyHash string) (bool, error) {
	if _, err := passlib.Verify(Config.Secret+":::"+username, accessKeyHash); err != nil {
		return false, errors.Wrap(err, "unable to verify access key")
	}
	if _, err := passlib.Verify(Config.Secret+":::"+accessKeyHash, secretKeyHash); err != nil {
		return false, errors.Wrap(err, "unable to verify secret key")
	}
	return true, nil
}
