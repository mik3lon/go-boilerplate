package user_application

type UserSignInQuery struct {
	Email    string
	Password string
}

func NewUserSignInQuery(email string, plainPassword string) *UserSignInQuery {
	return &UserSignInQuery{Email: email, Password: plainPassword}
}

func (r *UserSignInQuery) Id() string {
	return "user_sign_in_query"
}
