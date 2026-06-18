package mysql

import (
	"context"
	"database/sql"

	"github.com/kojima1128/portfolio-backend/internal/model"
)

// UserRepository is a MySQL implementation of repository.UserRepository.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new MySQL UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create inserts a new user into the database.
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	// TODO: Implement MySQL INSERT
	return nil
}

// FindByID retrieves a user by ID from the database.
func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	// TODO: Implement MySQL SELECT by ID
	return &model.User{
		ID:    id,
		Name:  "Sample User",
		Email: "user@example.com",
	}, nil
}

// FindAll retrieves all users from the database.
func (r *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	// TODO: Implement MySQL SELECT all
	return []*model.User{}, nil
}
