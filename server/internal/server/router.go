package server

import (
	"net/http"
)

type router struct {
	*http.ServeMux
	middlewares             []middleware
	prefix                  string
	handlerAfterMiddlewares *http.Handler
}

func newRouter(prefix string) *router {
	return &router{
		ServeMux:    http.NewServeMux(),
		middlewares: []middleware{},
		prefix:      prefix,
	}
}

func (r *router) Use(mw ...middleware) {
	r.middlewares = append(r.middlewares, mw...)
}

func mergeRoutes(routes ...*router) *http.ServeMux {
	h := http.NewServeMux()
	for _, r := range routes {
		h.Handle(r.prefix+"/", http.StripPrefix(r.prefix, *r.handlerAfterMiddlewares))
	}
	return h
}

type middleware func(http.Handler) http.Handler

func (r *router) applyMiddlewares() {
	r.handlerAfterMiddlewares = applyMiddlewares(r.middlewares, r.ServeMux)
}

func applyMiddlewares(mw []middleware, h http.Handler) *http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}
	return &h
}
