package kernel

import (
	"go-boilerplate/internal/app/module/user/application/register_user"
	user_application2 "go-boilerplate/internal/app/module/user/application/user_sign_in"
	"go-boilerplate/internal/app/module/user/infrastructure"
	user_ui "go-boilerplate/internal/app/module/user/ui"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/http/middleware"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib" // Import the pgx driver
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

	um.AddQuery(
		c.QueryBus,
		&user_application2.UserSignInQuery{},
		user_application2.NewUserSignInQueryHandler(r),
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

	c.Router.Handle(
		http.MethodPost,
		"/users/sign-in",
		user_ui.HandleUserSignIn(c.QueryBus),
	)
}
