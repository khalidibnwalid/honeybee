package resolvers_test

import (
	"khalidibnwalid/luma_server/internal/testutils"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	ctx := NewTestingGqlCtx(t)

	t.Run("should return the requested user data", func(t *testing.T) {
		user := testutils.MockUser(t, ctx.resolver.DB)

		// GraphQL query
		var resp struct {
			GetUser struct {
				ID        string
				Email     string
				Username  string
				AvatarURL *string
				CreatedAt *string
				UpdatedAt *string
			}
		}

		query := `
				query GetUser($id: ID!) {
					getUser(id: $id) {
						id
						email
						username
						avatarUrl
						createdAt
						updatedAt
					}
				}
				`

		// Execute
		ctx.client.Post(query, &resp, client.Var("id", user.ID.String()))

		assert.Equal(t, user.ID.String(), resp.GetUser.ID)
		assert.Equal(t, user.Email, resp.GetUser.Email)
		assert.Equal(t, user.AvatarURL.String, *resp.GetUser.AvatarURL)
		assert.NotNil(t, resp.GetUser.CreatedAt)
		assert.NotNil(t, resp.GetUser.UpdatedAt)
	})

	t.Run("should return an error if user not found", func(t *testing.T) {
		query := `
				query GetUser($id: ID!) {
					getUser(id: $id) {
						id
					}
				}`

		fakeID := uuid.New().String()
		err := ctx.client.Post(query, nil, client.Var("id", fakeID))

		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "USER_NOT_FOUND"))
		assert.True(t, strings.Contains(err.Error(), "getUser"))
	})

	t.Run("should return an error if user ID is invalid", func(t *testing.T) {
		query := `
				query GetUser($id: ID!) {
					getUser(id: $id) {
						id
					}
				}`

		err := ctx.client.Post(query, nil, client.Var("id", "INVALID_UUID"))

		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "INVALID_UUID_FORMAT"))
		assert.True(t, strings.Contains(err.Error(), "getUser"))
	})
}
