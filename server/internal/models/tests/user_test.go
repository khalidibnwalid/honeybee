package models_test

import (
	"khalidibnwalid/luma_server/internal/models"
	"khalidibnwalid/luma_server/internal/testutils"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	ctx := testutils.NewTestingServerHandlerCtx(t)

	t.Run("Create", func(t *testing.T) {
		user := models.User{
			Username:       "testuser",
			Email:          "test@example.com",
			HashedPassword: "hashedpassword",
		}

		err := user.Create(ctx.ServerContext.DB)

		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, user.ID)
		assert.NotZero(t, user.CreatedAt)
		assert.NotZero(t, user.UpdatedAt)
	})

	t.Run("GetByID", func(t *testing.T) {
		originalUser := models.User{
			Username:       "getbyiduser",
			Email:          "getbyid@example.com",
			HashedPassword: "hashedpassword",
		}

		err := originalUser.Create(ctx.ServerContext.DB)
		assert.NoError(t, err)

		// retrieve the user by ID
		var retrievedUser models.User
		err = retrievedUser.GetByID(ctx.ServerContext.DB, originalUser.ID)

		assert.NoError(t, err)
		assert.Equal(t, originalUser.ID, retrievedUser.ID)
		assert.Equal(t, originalUser.Username, retrievedUser.Username)
		assert.Equal(t, originalUser.Email, retrievedUser.Email)
	})

	t.Run("GetByUsername", func(t *testing.T) {
		originalUser := models.User{
			Username:       "getbyusername",
			Email:          "getbyusername@example.com",
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
		user := models.User{
			Username:       "updateuser",
			Email:          "update@example.com",
			HashedPassword: "hashedpassword",
		}
		err := user.Create(ctx.ServerContext.DB)
		assert.NoError(t, err)

		// Update user fields
		user.Username = "updateduservalue"
		user.Email = "updated@example.com"

		err = user.Update(ctx.ServerContext.DB)
		assert.NoError(t, err)

		// Retrieve the user to confirm updates
		var retrievedUser models.User
		err = retrievedUser.GetByID(ctx.ServerContext.DB, user.ID)

		assert.NoError(t, err)
		assert.Equal(t, "updateduservalue", retrievedUser.Username)
		assert.Equal(t, "updated@example.com", retrievedUser.Email)
	})

	t.Run("Delete", func(t *testing.T) {
		user := models.User{
			Username:       "deleteuser",
			Email:          "delete@example.com",
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
