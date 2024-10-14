package user_application

type RegisterUserCommand struct {
	userId        string
	username      string
	email         string
	plainPassword string
	name          string
	surname       string
}

func NewRegisterUserCommand(
	userId string,
	username string,
	email string,
	plainPassword string,
	name string,
	surname string,
) *RegisterUserCommand {
	return &RegisterUserCommand{userId: userId, username: username, email: email, plainPassword: plainPassword, name: name, surname: surname}
}

func (r *RegisterUserCommand) Id() string {
	return "register_user_command"
}
