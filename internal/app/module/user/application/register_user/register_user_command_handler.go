package user_application

import (
	"context"
	"errors"
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"go-boilerplate/internal/pkg/domain/bus"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserCommandHandler struct {
	r user_domain.UserRepository
}

func NewRegisterUserCommandHandler(r user_domain.UserRepository) *RegisterUserCommandHandler {
	return &RegisterUserCommandHandler{r: r}
}

func (ruqh *RegisterUserCommandHandler) Handle(ctx context.Context, command bus.Command) error {
	ruq, ok := command.(*RegisterUserCommand)
	if !ok {
		return bus.NewCommandNotValid("command not valid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ruq.plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	u := user_domain.CreateUser(
		ruq.userId,
		ruq.username,
		ruq.email,
		string(hashedPassword), // Store hashed password (salt is part of this)
		ruq.name,
		ruq.surname,
		"ROLE_TEMP",
	)

	return ruqh.r.Save(ctx, u)
}
