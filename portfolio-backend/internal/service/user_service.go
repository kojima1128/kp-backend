package service

import (
	"context"

	"github.com/kojima1128/portfolio-backend/internal/model"
	"github.com/kojima1128/portfolio-backend/internal/repository"
)

// UserService はユーザーに関するビジネスロジックを担当します。
// API・ワーカー双方から共通で利用されます。
type UserService struct {
	repo repository.UserRepository
}

// NewUserService は UserService を生成します。
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetUser は指定した ID のユーザーを取得します。
func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}

// GetUsers は全ユーザーを取得します。
func (s *UserService) GetUsers(ctx context.Context) ([]*model.User, error) {
	return s.repo.FindAll(ctx)
}

// CreateUser は新しいユーザーを作成します。
func (s *UserService) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return s.repo.Create(ctx, input)
}
