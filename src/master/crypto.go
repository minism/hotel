package master

import (
	"crypto/rand"
	"encoding/base32"
)

func GenerateRandomBytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func GenerateRandomB32String(length int) (string, error) {
	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	buf, err := GenerateRandomBytes(length)
	return encoder.EncodeToString(buf), err
}
