package createHash

import (
	"crypto/rand"
	"encoding/hex"
)

func CreateHash() (string, error) {
	hash := make([]byte, 16)
	_, err := rand.Read(hash)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash), nil
}