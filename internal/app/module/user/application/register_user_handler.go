package user_application

import (
	"context"
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"go-boilerplate/internal/pkg/domain/bus"
)

type RegisterUserCommand struct {
	userId   string
	username string
	email    string
}

func NewRegisterUserCommand(userId string, username string, email string) *RegisterUserCommand {
	return &RegisterUserCommand{userId: userId, username: username, email: email}
}

func (r *RegisterUserCommand) Id() string {
	return "register_user_query"
}

type RegisterUserCommandHandler struct {
	r user_domain.UserRepository
}

func NewRegisterUserCommandHandler(r user_domain.UserRepository) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{r: r}
}

func (ruqh *RegisterUserCommandHandler) Handle(_ context.Context, command bus.Command) error {
	ruq, ok := command.(*RegisterUserCommand)
	if !ok {
		return bus.NewCommandNotValid("command not valid")
	}

	u := user_domain.CreateUser(
		ruq.userId,
		ruq.username,
		ruq.email,
		"",
		"",
		"",
		"",
		"",
	)

	return ruqh.r.Save(u)
}
