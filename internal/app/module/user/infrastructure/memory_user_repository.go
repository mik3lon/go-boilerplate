package user_infrastructure

import (
	user_domain "go-boilerplate/internal/app/module/user/domain"
)

// MemoryUserRepository is an in-memory implementation of UserRepository.
type MemoryUserRepository struct {
	users map[string]*user_domain.User
}

// NewMemoryUserRepository initializes a new in-memory repository.
func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{users: make(map[string]*user_domain.User)}
}

// Save stores a user in memory.
func (r *MemoryUserRepository) Save(user *user_domain.User) error {
	r.users[user.ID] = user
	return nil
}

// FindByID retrieves a user by ID.
func (r *MemoryUserRepository) FindByID(id string) (*user_domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, user_domain.ErrUserNotFound
	}
	return user, nil
}
