package models_test

import (
	"khalidibnwalid/luma_server/internal/crypto"
	"khalidibnwalid/luma_server/internal/models"
	"khalidibnwalid/luma_server/internal/testutils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	ctx := testutils.NewTestingServerHandlerCtx(t)

	t.Run(("SetPassword"), func(t *testing.T) {
		password := crypto.RandomString(10)
		user := models.User{
			Username: crypto.RandomString(10),
			Email:    crypto.RandomString(10) + "@example.com",
		}
		user.SetPassword(password)

		assert.NotEqual(t, user.HashedPassword, password, "Hashed password should not be equal to the plain password")
		assert.NotEmpty(t, user.HashedPassword, "Hashed password should not be empty")

		assert.NoError(t, crypto.VerifyHashWithSalt(password, user.HashedPassword), "Hashed password verification failed")
	})

	t.Run("VerifyPassword", func(t *testing.T) {
		rawPassword := crypto.RandomString(10)
		user := models.User{
			Username: crypto.RandomString(10),
			Email:    crypto.RandomString(10) + "@example.com",
		}
		user.SetPassword(rawPassword)

		t.Run("should verify correct password", func(t *testing.T) {
			err := user.VerifyPassword(rawPassword)
			assert.NoError(t, err, "Password verification failed")
		})

		t.Run("should return error for wrong password", func(t *testing.T) {
			wrongPassword := crypto.RandomString(12) // intentionally different length => 12 != 10
			err := user.VerifyPassword(wrongPassword)
			assert.Error(t, err, "Expected error for wrong password")
		})
	})

	t.Run("Create & GetByID", func(t *testing.T) {
		username := crypto.RandomString(10)
		user := models.User{
			Username:       username,
			Email:          username + "@example.com",
			HashedPassword: "hashedpassword",
		}

		err := user.Create(ctx.ServerContext.DB)

		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, user.ID)
		assert.NotZero(t, user.CreatedAt)
		assert.NotZero(t, user.UpdatedAt)

		// Check if the user exists in the database
		var retrievedUser models.User
		err = retrievedUser.GetByID(ctx.ServerContext.DB, user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user.ID, retrievedUser.ID)
		assert.Equal(t, user.Username, retrievedUser.Username)
	})

	t.Run("GetByUsername", func(t *testing.T) {
		username := crypto.RandomString(10)

		originalUser := models.User{
			Username:       username,
			Email:          username + "@example.com",
			HashedPassword: "hashedpassword",
		}
		err := originalUser.Create(ctx.ServerContext.DB)
		assert.NoError(t, err)

		// retrieve the user by username
		var retrievedUser models.User
		err = retrievedUser.GetByUsername(ctx.ServerContext.DB, originalUser.Username)

		assert.NoError(t, err)
		assert.Equal(t, originalUser.ID, retrievedUser.ID)
		assert.Equal(t, originalUser.Username, retrievedUser.Username)
		assert.Equal(t, originalUser.Email, retrievedUser.Email)
	})

	t.Run("Update", func(t *testing.T) {
		originalUsername := crypto.RandomString(10)

		user := models.User{
			Username:       originalUsername,
			Email:          originalUsername + "@example.com",
			HashedPassword: "hashedpassword",
		}
		err := user.Create(ctx.ServerContext.DB)
		assert.NoError(t, err)

		// Update user fields
		user.Username = crypto.RandomString(10)

		err = user.Update(ctx.ServerContext.DB)
		assert.NoError(t, err)

		// Retrieve the user to confirm updates
		var retrievedUser models.User
		err = retrievedUser.GetByID(ctx.ServerContext.DB, user.ID)

		assert.NoError(t, err)
		assert.Equal(t, user.Username, retrievedUser.Username)
		// Ensure the Email and ID remains the same
		assert.Equal(t, user.ID, retrievedUser.ID)
		assert.Equal(t, user.Email, retrievedUser.Email)
	})

	t.Run("Delete", func(t *testing.T) {
		username := crypto.RandomString(10)
		user := models.User{
			Username:       username,
			Email:          username + "@example.com",
			HashedPassword: "hashedpassword",
		}
		err := user.Create(ctx.ServerContext.DB)
		assert.NoError(t, err)

		err = user.Delete(ctx.ServerContext.DB)
		assert.NoError(t, err)

		// Try to retrieve the deleted user, should return error
		var retrievedUser models.User
		err = retrievedUser.GetByID(ctx.ServerContext.DB, user.ID)
		assert.Error(t, err)
	})
}
