package graph

import "github.com/kojima1128/portfolio-backend/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver はGraphQLリゾルバーのルート構造体です。
// サービス層への依存を保持します。
type Resolver struct {
	UserService *service.UserService
}
