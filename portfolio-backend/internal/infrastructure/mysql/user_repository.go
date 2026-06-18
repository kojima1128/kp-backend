package mysql

import (
	"context"
	"database/sql"

	"github.com/kojima1128/portfolio-backend/internal/model"
)

// UserRepository は MySQL を用いた UserRepository の実装です。
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository は UserRepository を生成します。
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByID は指定した ID のユーザーを取得します。
func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	// TODO: implement
	return nil, nil
}

// FindAll は全ユーザーを取得します。
func (r *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	// TODO: implement
	return []*model.User{}, nil
}

// Create は新しいユーザーを作成します。
func (r *UserRepository) Create(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	// TODO: implement
	return nil, nil
}
