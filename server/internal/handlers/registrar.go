package handlers

import (
	"khalidibnwalid/luma_server/internal/middlewares"
	"khalidibnwalid/luma_server/internal/server"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/joho/godotenv/autoload"
)

// Register the routes for the server here
func (s *ServerHandlerContext) registrar() *http.ServeMux {
	h := http.NewServeMux()
	r := server.NewRouter("/")

	// r.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("1"))
	// 		next.ServeHTTP(w, r)
	// 		w.Write([]byte("1"))
	// 	})
	// })

	r.Use(middlewares.CORS)

	if strings.ToLower(os.Getenv("PRODUCTION")) != "true" {
		r.Handle("/graphpg", playground.Handler("GraphQL playground", "/query"))
	}

	r.Handle("/query", s.GraphQLHandler())

	handler := r.ApplyMiddlewares()

	h.Handle("/", *handler)

	return h
}
