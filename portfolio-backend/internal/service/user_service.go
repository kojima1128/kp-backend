package service

import (
	"context"

	"github.com/kojima1128/portfolio-backend/internal/model"
	"github.com/kojima1128/portfolio-backend/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	user := &model.User{Name: input.Name, Email: input.Email}
	return user, s.repo.Create(ctx, user)
}
