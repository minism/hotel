package main

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomBytes(length int) ([]byte, error) {
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func GenerateRandomB64String(length int) (string, error) {
	buf, err := GenerateRandomBytes(length)
	return base64.URLEncoding.EncodeToString(buf), err
}
