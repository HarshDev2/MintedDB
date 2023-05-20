package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateID() string {
	randomBytes := make([]byte, 16) // 16 bytes = 128 bits
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err) // handle error appropriately in your code
	}

	hash := sha256.Sum256(randomBytes)
	return hex.EncodeToString(hash[:])
}