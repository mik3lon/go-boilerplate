package user_ui

import (
	"context"
	"encoding/json"
	"errors"
	user_application "go-boilerplate/internal/app/module/user/application/register_user"
	"go-boilerplate/internal/pkg/domain/bus"
	http_writer "go-boilerplate/pkg/http/writer"
	"net/http"
)

func HandleRegisterUser(cb bus.CommandBus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
			Name     string `json:"name"`
			Surname  string `json:"surname"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http_writer.WriteError(w, errors.New("invalid input"), http.StatusBadRequest)
			return
		}

		err := cb.Dispatch(
			context.Background(),
			user_application.NewRegisterUserCommand(
				req.ID,
				req.Username,
				req.Email,
				req.Password,
				req.Name,
				req.Surname,
			),
		)

		if err != nil {
			http_writer.WriteError(w, err, http.StatusBadRequest)
			return
		}

		http_writer.WriteNoContent(w)
	}
}
