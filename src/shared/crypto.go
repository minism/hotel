package shared

import (
	"crypto/rand"
	"encoding/base32"
)

// GenerateRandomBytes returns an array of `length` random bytes.
func GenerateRandomBytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// GenerateRandomB32String returns a base32-encoded string of `length` random bytes.
func GenerateRandomB32String(length int) (string, error) {
	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	buf, err := GenerateRandomBytes(length)
	return encoder.EncodeToString(buf), err
}
