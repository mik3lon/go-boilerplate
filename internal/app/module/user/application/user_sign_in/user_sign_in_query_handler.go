package user_application

import (
	"context"
	"errors"
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"go-boilerplate/internal/pkg/domain/bus"
	"golang.org/x/crypto/bcrypt"
)

type UserSignInQueryHandler struct {
	r user_domain.UserRepository
}

func NewUserSignInQueryHandler(r user_domain.UserRepository) *UserSignInQueryHandler {
	return &UserSignInQueryHandler{r: r}
}

func (siqh *UserSignInQueryHandler) Handle(ctx context.Context, query bus.Query) (interface{}, error) {
	siq, ok := query.(*UserSignInQuery)
	if !ok {
		return nil, bus.NewInvalidQuery(query)
	}

	// Retrieve the user by username from the repository
	u, err := siqh.r.FindByEmail(ctx, siq.Email)
	if err != nil {
		// Handle the case where the user is not found
		return nil, errors.New("user not found")
	}

	// Compare the provided plain password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(u.HashedPassword()), []byte(siq.Password))
	if err != nil {
		// Password doesn't match
		return nil, errors.New("invalid credentials")
	}

	return u, nil
}
