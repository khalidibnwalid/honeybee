package testutils

import (
	"khalidibnwalid/luma_server/internal/crypto"
	"khalidibnwalid/luma_server/internal/database"
	"khalidibnwalid/luma_server/internal/models"
	"testing"
)

func MockUser(t *testing.T, Db *database.Database) *models.User {
	t.Helper()

	username := crypto.RandomString(10)
	user := &models.User{
		Username: username,
		Email:    username + "@example.com",
	}

	user.SetPassword("password" + username)
	if err := user.Create(Db); err != nil {
		t.Fatalf("Failed to create mock user: %v", err)
	}

	return user
}
