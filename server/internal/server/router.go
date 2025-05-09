package server

import (
	"net/http"
)

type Router struct {
	*http.ServeMux
	middlewares             []middleware
	prefix                  string
	handlerAfterMiddlewares *http.Handler
}

func NewRouter(prefix string) *Router {
	return &Router{
		ServeMux:    http.NewServeMux(),
		middlewares: []middleware{},
		prefix:      prefix,
	}
}

func (r *Router) Use(mw ...middleware) {
	r.middlewares = append(r.middlewares, mw...)
}

func MergeRoutes(routes ...*Router) *http.ServeMux {
	h := http.NewServeMux()
	for _, r := range routes {
		h.Handle(r.prefix+"/", http.StripPrefix(r.prefix, *r.handlerAfterMiddlewares))
	}
	return h
}

type middleware func(http.Handler) http.Handler

func (r *Router) ApplyMiddlewares() {
	r.handlerAfterMiddlewares = ApplyMiddlewares(r.middlewares, r.ServeMux)
}

func ApplyMiddlewares(mw []middleware, h http.Handler) *http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}
	return &h
}
