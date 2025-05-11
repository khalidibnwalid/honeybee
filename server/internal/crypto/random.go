package crypto

import (
	"crypto/rand"
	"encoding/base64"
)

// Cryptographically secure random bytes generator
func RandomBytes(length uint32) []byte {
	buffer := make([]byte, length)
	rand.Read(buffer)
	return buffer
}

func RandomString(length uint32) string {
	bytes := RandomBytes(length)
	return base64.RawStdEncoding.EncodeToString(bytes)
}
