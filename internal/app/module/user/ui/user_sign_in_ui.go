package user_ui

import (
	"context"
	"encoding/json"
	"errors"
	user_application "go-boilerplate/internal/app/module/user/application/user_sign_in"
	"go-boilerplate/internal/pkg/domain/bus"
	http_writer "go-boilerplate/pkg/http/writer"
	"net/http"
)

func HandleUserSignIn(qb bus.QueryBus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email         string `json:"email"`
			PlainPassword string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http_writer.WriteError(w, errors.New("invalid input"), http.StatusBadRequest)
			return
		}

		user, err := qb.Ask(
			context.Background(),
			user_application.NewUserSignInQuery(req.Email, req.PlainPassword),
		)

		switch err {
		case nil:
			http_writer.WriteJSON(w, user, http.StatusOK)
			return
		default:
			http_writer.WriteError(w, err, http.StatusInternalServerError)
			return
		}
	}
}
