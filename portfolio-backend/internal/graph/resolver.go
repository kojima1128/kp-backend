package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/kojima1128/portfolio-backend/internal/model"
	"github.com/kojima1128/portfolio-backend/internal/service"
)

type Resolver struct {
	userService *service.UserService
}

func NewResolver(userService *service.UserService) *Resolver {
	return &Resolver{userService: userService}
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() QueryResolver       { return &queryResolver{r} }

var _ graphql.ExecutableSchema = (*executableSchema)(nil)

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

type executableSchema struct{}

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	return r.userService.GetUser(ctx, id)
}
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return r.userService.ListUsers(ctx)
}
func (r *mutationResolver) CreateUser(ctx context.Context, input CreateUserInput) (*User, error) {
	return r.userService.CreateUser(ctx, model.CreateUserInput(input))
}
