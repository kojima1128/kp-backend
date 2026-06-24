package repository

import (
	"context"

	"github.com/kojima1128/portfolio-backend/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type inMemoryUserRepository struct{}

func NewUserRepository() UserRepository {
	return &inMemoryUserRepository{}
}

func (r *inMemoryUserRepository) Create(ctx context.Context, user *model.User) error { return nil }
func (r *inMemoryUserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	return &model.User{ID: id, Name: "Sample User", Email: "user@example.com"}, nil
}
func (r *inMemoryUserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}
func (r *inMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}
func (r *inMemoryUserRepository) Update(ctx context.Context, user *model.User) error { return nil }
func (r *inMemoryUserRepository) Delete(ctx context.Context, id string) error        { return nil }
