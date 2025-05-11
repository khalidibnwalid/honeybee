package testutils

import (
	"khalidibnwalid/luma_server/internal/database"
	"khalidibnwalid/luma_server/internal/handlers"
	"khalidibnwalid/luma_server/internal/server"
	"testing"
)

const uri = "postgres://admin:123qweasd@localhost:5432/testingLuma?sslmode=disable&TimeZone=UTC"

func NewTestingServerHandlerCtx(t *testing.T) *handlers.ServerHandlerContext {
	t.Helper()

	db, err := database.NewClient(uri)
	if err != nil {
		t.Fatalf("Failed to create database client: %v", err)
	}

	tx := db.Client.Begin()
	tx.SavePoint("before_test")

	DB := &database.Database{
		Client: tx,
	}

	t.Cleanup(func() {
		tx.RollbackTo("before_test")
	})

	serverHandlerContext := &handlers.ServerHandlerContext{
		ServerContext: &server.ServerContext{
			DB: DB,
		},
	}

	return serverHandlerContext
}
