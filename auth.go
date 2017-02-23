package auth

import passlib "gopkg.in/hlandau/passlib.v1"

func Hash(username string) (accessKeyHash string, secretKeyHash string, err error) {
	accessKeyHash, err = passlib.Hash(username + Config.Secret)
	if err != nil {
		return
	}
	secretKeyHash, err = passlib.Hash(accessKeyHash + Config.Secret)
	return
}

func Verify(username, accessKeyHash, secretKeyHash string) bool {
	if _, err := passlib.Verify(accessKeyHash, username+Config.Secret); err != nil {
		return false
	}
	if _, err := passlib.Verify(secretKeyHash, accessKeyHash+Config.Secret); err != nil {
		return false
	}
	return true
}
