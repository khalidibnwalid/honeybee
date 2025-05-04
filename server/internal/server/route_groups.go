package server

import (
	"net/http"
)

// Register the routes for the server here
func registeredRoutes() *http.ServeMux {
	routes := mergeRoutes(
		unauthRoutes("/v1/login"),
		authRoutes("/v1"),
	)
	return routes
}

func authRoutes(prefix string) *router {
	r := newRouter(prefix)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("1"))
			next.ServeHTTP(w, r)
			w.Write([]byte("1"))
		})
	})

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("2"))
			next.ServeHTTP(w, r)
			w.Write([]byte("2"))
		})
	})

	// Register the routes for the server here
	r.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Worldd!"))

	})

	r.applyMiddlewares()
	return r
}

func unauthRoutes(prefix string) *router {
	r := newRouter(prefix)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("3"))
			next.ServeHTTP(w, r)
			w.Write([]byte("3"))
		})
	})

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("4"))
			next.ServeHTTP(w, r)
			w.Write([]byte("4"))
		})
	})

	// Register the routes for the server here
	r.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Worldd!"))
	})

	r.applyMiddlewares()
	return r
}
