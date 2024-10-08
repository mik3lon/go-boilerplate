package user_infrastructure

import (
	user_domain2 "go-boilerplate/internal/app/module/user/domain"
)

// MemoryUserRepository is an in-memory implementation of UserRepository.
type MemoryUserRepository struct {
	users map[string]*user_domain2.User
}

// NewMemoryUserRepository initializes a new in-memory repository.
func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{users: make(map[string]*user_domain2.User)}
}

// Save stores a user in memory.
func (r *MemoryUserRepository) Save(user *user_domain2.User) error {
	r.users[user.ID] = user
	return nil
}

// FindByID retrieves a user by ID.
func (r *MemoryUserRepository) FindByID(id string) (*user_domain2.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, user_domain2.ErrUserNotFound
	}
	return user, nil
}
