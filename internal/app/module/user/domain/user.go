package user_domain

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;size:255"`
	Email     string    `json:"email" gorm:"uniqueIndex;size:255"`
	Password  string    `json:"-" gorm:"size:255"`
	Salt      string    `json:"-" gorm:"size:255"`
	Name      string    `json:"name" gorm:"size:255"`
	Surname   string    `json:"surname" gorm:"size:255"`
	Role      string    `json:"role" gorm:"size:100"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// CreateUser creates a new User entity.
func CreateUser(id, username, email, password, salt, name, surname, role string) *User {
	return &User{
		ID:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		Salt:      salt,
		Name:      name,
		Surname:   surname,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
