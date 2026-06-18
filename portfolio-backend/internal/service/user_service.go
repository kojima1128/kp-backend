package service

import (
	"context"

	"github.com/kojima1128/portfolio-backend/internal/model"
	"github.com/kojima1128/portfolio-backend/internal/repository"
)

// UserService handles business logic for users.
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser retrieves a user by ID.
func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}

// GetUsers retrieves all users.
func (s *UserService) GetUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.FindAll(ctx)
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
