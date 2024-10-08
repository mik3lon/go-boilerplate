package user_domain

type UserRepository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
}
