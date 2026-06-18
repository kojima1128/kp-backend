package repository

import (
	"context"

	"github.com/kojima1128/portfolio-backend/internal/model"
)

// UserRepository はユーザーの永続化処理を定義するインターフェースです。
// API・ワーカー双方から共通で利用されます。
type UserRepository interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	Create(ctx context.Context, input model.CreateUserInput) (*model.User, error)
}
