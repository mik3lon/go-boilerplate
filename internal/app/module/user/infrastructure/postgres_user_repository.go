package user_infrastructure

import (
	"errors"
	user_domain "go-boilerplate/internal/app/module/user/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresUserRepository is a Postgres implementation of UserRepository using Gorm.
type PostgresUserRepository struct {
	DB *gorm.DB
}

// NewPostgresUserRepository initializes a new Postgres user repository.
func NewPostgresUserRepository(dsn string) (*PostgresUserRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Ensure the User table exists
	err = db.AutoMigrate(&user_domain.User{})
	if err != nil {
		return nil, err
	}

	return &PostgresUserRepository{
		DB: db,
	}, nil
}

// Save stores a user in the Postgres database.
func (r *PostgresUserRepository) Save(user *user_domain.User) error {
	result := r.DB.Save(user)
	return result.Error
}

// FindByID retrieves a user by ID from the Postgres database.
func (r *PostgresUserRepository) FindByID(id string) (*user_domain.User, error) {
	var user user_domain.User
	result := r.DB.First(&user, "id = ?", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, user_domain.ErrUserNotFound
	}
	return &user, result.Error
}
