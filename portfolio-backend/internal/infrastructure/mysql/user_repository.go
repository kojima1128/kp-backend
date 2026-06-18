package mysql

import "github.com/kojima1128/portfolio-backend/internal/repository"

type UserRepository struct{}

func NewUserRepository() repository.UserRepository {
	return repository.NewUserRepository()
}
