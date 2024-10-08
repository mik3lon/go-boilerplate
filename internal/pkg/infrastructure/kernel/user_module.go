package kernel

import (
	user_application "go-boilerplate/internal/app/module/user/application"
	"go-boilerplate/internal/app/module/user/infrastructure"
	user_ui "go-boilerplate/internal/app/module/user/ui"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/http/middleware"
	"net/http"
)

type UserModule struct {
	CQRSModule
}

// InitUserModule creates a new instance of UserModule.
func InitUserModule(c *Kernel, cnf *config.Config) *UserModule {
	r, err := user_infrastructure.NewPostgresUserRepository(cnf.DatabaseDSN)
	if err != nil {
		panic(err)
	}

	um := &UserModule{}
	um.AddCommand(
		c.CommandBus,
		&user_application.RegisterUserCommand{},
		user_application.NewRegisterUserCommandHandler(r),
	)

	return um
}

// RegisterRoutes registers the user module routes.
func (m *UserModule) RegisterRoutes(c *Kernel) {
	c.Router.WithMiddleware(middleware.LoggingMiddleware).Handle(
		http.MethodPost,
		"/users",
		user_ui.HandleRegisterUser(c.CommandBus),
	)
}
