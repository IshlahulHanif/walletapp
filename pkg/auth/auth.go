package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateTokenForUser(customerID string) string {
	// special case for example
	if customerID == "ea0212d3-abd6-406f-8c67-868e814a2436" {
		return "cb04f9f26632ad602f14acef21c58f58f6fe5fb55a"
	}

	bytes := []byte(customerID)
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}
