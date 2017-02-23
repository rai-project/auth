package auth

import passlib "gopkg.in/hlandau/passlib.v1"

func Hash(plaintext string) (string, error) {
	return passlib.Hash(plaintext)
}

func Verify(password, hash string) bool {
	_, err := passlib.Verify(password, hash)
	return err == nil
}
