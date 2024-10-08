package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinRouter is an adapter for the gin router.
type GinRouter struct {
	Router     *gin.Engine
	middleware []Middleware
}

// NewGinRouter creates a new instance of GinRouter.
func NewGinRouter() *GinRouter {
	return &GinRouter{
		Router: gin.Default(),
	}
}

// WithMiddleware adds middleware to the router and returns the updated router.
func (g *GinRouter) WithMiddleware(middleware ...Middleware) Router {
	g.middleware = append(g.middleware, middleware...)
	return g
}

// Handle registers a new route and applies the middleware to the handler.
func (g *GinRouter) Handle(method, path string, handler http.HandlerFunc) {
	// Apply the middleware chain to the handler
	finalHandler := handler
	for _, m := range g.middleware {
		finalHandler = m(finalHandler)
	}

	switch method {
	case http.MethodGet:
		g.Router.GET(path, gin.WrapH(finalHandler))
	case http.MethodPost:
		g.Router.POST(path, gin.WrapH(finalHandler))
	}
}

// Serve starts the Gin server at the given address.
func (g *GinRouter) Serve(addr string) error {
	return g.Router.Run(addr)
}
