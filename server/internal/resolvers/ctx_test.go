// this is testing utlity, not an actual test
package resolvers_test

import (
	"khalidibnwalid/luma_server/internal/database"
	"khalidibnwalid/luma_server/internal/graph"
	"khalidibnwalid/luma_server/internal/resolvers"
	"khalidibnwalid/luma_server/internal/server"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

const uri = "postgres://admin:123qweasd@localhost:5432/testingLuma?sslmode=disable&TimeZone=UTC"

type TestingGQLCtx struct {
	resolver *resolvers.Resolver
	client   *client.Client
}

func NewTestingGqlCtx(t *testing.T) *TestingGQLCtx {
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

	res := &resolvers.Resolver{
		ServerContext: &server.ServerContext{
			DB: DB,
		},
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: res}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	c := client.New(srv)

	return &TestingGQLCtx{
		resolver: res,
		client:   c,
	}
}
