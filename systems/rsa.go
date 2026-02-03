package systems

import (
	"crypto/rand"
	"crypto/rsa"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func InitRSA() error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	PrivateKey = key
	PublicKey = &key.PublicKey

	return nil
}
