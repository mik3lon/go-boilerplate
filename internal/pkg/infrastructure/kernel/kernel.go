package kernel

import (
	"context"
	"go-boilerplate/internal/pkg/domain/bus"
	bus_infra "go-boilerplate/internal/pkg/infrastructure/bus"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/router"
	"net/http"
)

type Kernel struct {
	Router     router.Router
	Modules    Modules
	QueryBus   bus.QueryBus
	CommandBus bus.CommandBus
	server     *http.Server
}

// Init initializes the container with a router implementation.
func Init(cnf *config.Config) *Kernel {
	r := router.NewGinRouter()

	k := &Kernel{
		Router:     r,
		QueryBus:   bus_infra.InitQueryBus(),
		CommandBus: bus_infra.InitCommandBus(),
		server: &http.Server{
			Addr:    cnf.AddressPort,
			Handler: r.Handler(),
		},
	}

	userModule := InitUserModule(k, cnf)

	k.addModule(userModule)

	k.RegisterModuleRoutes()

	return k
}

// RegisterModuleRoutes allows each module to register its routes.
func (c *Kernel) RegisterModuleRoutes() {
	for _, m := range c.Modules {
		m.RegisterRoutes(c)
	}
}

// StartServer starts the HTTP server.
func (c *Kernel) StartServer() error {
	return c.server.ListenAndServe()
}

func (c *Kernel) addModule(module Module) {
	c.Modules = append(c.Modules, module)
}

func (c *Kernel) ShutdownServer(ctx context.Context) error {
	return c.server.Shutdown(ctx)
}
