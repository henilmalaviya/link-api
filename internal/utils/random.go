package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	// Encode to Base64 URL encoding and trim to requested length
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}
