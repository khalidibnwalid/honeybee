package handlers

import (
	"khalidibnwalid/luma_server/internal/server"
	"net/http"
)

// Register the routes for the server here
func (s *ServerHandlerContext) registrar() *http.ServeMux {
	routes := server.MergeRoutes(
		// s.unauthRoutes("/v1/login"),
		s.authRoutes("/v1"),
	)
	return routes
}

func (s *ServerHandlerContext) authRoutes(prefix string) *server.Router {
	r := server.NewRouter(prefix)

	// r.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.Write([]byte("1"))
	// 		next.ServeHTTP(w, r)
	// 		w.Write([]byte("1"))
	// 	})
	// })

	// Register the routes for the server here
	// r.HandleFunc("GET /hello",s.PostUser)

	r.ApplyMiddlewares()
	return r
}

// func (s *ServerHandlerContext) unauthRoutes(prefix string) *router {
// 	r := newRouter(prefix)

// 	// Register the routes for the server here
// 	r.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello, Worldd!"))
// 	})

// 	r.applyMiddlewares()
// 	return r
// }
