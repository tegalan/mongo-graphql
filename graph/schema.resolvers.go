package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"mongo-graph/graph/generated"
	"mongo-graph/graph/model"
	"mongo-graph/services/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*user.User, error) {
	u := user.User{
		Name:  input.Name,
		Email: input.Email,
	}
	r.UserService.Create(&u)
	return &u, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*user.User, error) {
	users := r.UserService.All()
	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.UserService.FindByID(id), nil
}

func (r *userResolver) ID(ctx context.Context, obj *user.User) (string, error) {
	return obj.ID.Hex(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
