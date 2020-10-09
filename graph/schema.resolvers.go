package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"mongo-graph/graph/generated"
	"mongo-graph/graph/model"
	"mongo-graph/services/post"
	"mongo-graph/services/user"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*user.User, error) {
	u := user.User{
		Name:  input.Name,
		Email: input.Email,
	}
	r.UserService.Create(&u)
	return &u, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*post.Post, error) {
	uid, _ := primitive.ObjectIDFromHex(input.UserID)
	p := post.Post{
		Title:  input.Title,
		Body:   input.Body,
		UserID: uid,
	}
	err := r.PostService.Create(&p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *postResolver) ID(ctx context.Context, obj *post.Post) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*user.User, error) {
	return r.UserService.All()
}

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	return r.UserService.FindByID(id)
}

func (r *queryResolver) Posts(ctx context.Context) ([]*post.Post, error) {
	return r.PostService.All()
}

func (r *userResolver) ID(ctx context.Context, obj *user.User) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *userResolver) Posts(ctx context.Context, obj *user.User) ([]*post.Post, error) {
	return r.PostService.FindByUser(obj.ID.Hex())
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Post(ctx context.Context, id string) (*post.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
