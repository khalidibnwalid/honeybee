package crypto_test

import (
	"fmt"
	"khalidibnwalid/luma_server/internal/crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeDeserializeHash(t *testing.T) {
	t.Run("should serialize and deserialize a hash with salt", func(t *testing.T) {
		valueToHash := crypto.RandomString(10)
		// original hash
		hash, salt := crypto.HashWithSalt(valueToHash)
		serialized := crypto.SerializeHashWithSalt(hash, salt)

		// Test deserialization
		deserializedHash, deserializedSalt, err := crypto.DeserializeHash(serialized)
		assert.Nil(t, err, fmt.Sprintf("Failed to deserialize hash: %v", err))

		assert.Equal(t, hash, deserializedHash, "Original and deserialized hash are not equal")
		assert.Equal(t, salt, deserializedSalt, "Original and deserialized salt are not equal")

	})

	t.Run("should return error if invalid serialized format was provided", func(t *testing.T) {
		_, _, err := crypto.DeserializeHash("invalid$format")
		assert.Equal(t, err, crypto.ErrInvalidHashFormat, fmt.Sprintf("Expected ErrInvalidHashFormat, got %v", err))
	})
}

func TestVerifyHashWithSalt(t *testing.T) {
	securePassword := crypto.RandomString(10)

	hash, salt := crypto.HashWithSalt(securePassword)
	serialized := crypto.SerializeHashWithSalt(hash, salt)

	t.Run("should verify a hash with salt", func(t *testing.T) {

		err := crypto.VerifyHashWithSalt(securePassword, serialized)
		assert.Nil(t, err, fmt.Sprintf("Failed to verify hash: %v", err))
	})

	t.Run("should return error if hash verification failed", func(t *testing.T) {
		wrongPassword := "wrongPassword"
		err := crypto.VerifyHashWithSalt(wrongPassword, serialized)
		if err != crypto.ErrHashVerificationFailed {
			t.Errorf("Expected ErrHashVerificationFailed, got %v", err)
		}
	})

	t.Run("should return error if invalid serialized format was provided", func(t *testing.T) {
		err := crypto.VerifyHashWithSalt(securePassword, "invalid$format")
		if err != crypto.ErrInvalidHashFormat {
			t.Errorf("Expected ErrInvalidHashFormat, got %v", err)
		}
	})
}
