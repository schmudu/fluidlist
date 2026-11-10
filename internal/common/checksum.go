package common

import (
	"crypto/sha256"
	"fmt"
)

func GetChecksum(content []byte) string {
	// Convert string to bytes and compute the SHA256 checksum
	hash := sha256.Sum256(content)

	return fmt.Sprintf("%x", hash)
}
