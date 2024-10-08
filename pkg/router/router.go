package router

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Router abstracts the HTTP router functionality.
type Router interface {
	Handle(method, path string, handler http.HandlerFunc)
	WithMiddleware(middleware ...Middleware) Router
	Serve(addr string) error
}
