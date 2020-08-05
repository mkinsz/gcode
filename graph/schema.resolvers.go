package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gcode/graph/generated"
	"gcode/graph/model"
	"strconv"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	u := model.User{ID: "3", Name: "Andy"}
	r.users = append(r.users, u)
	return &u, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUsers(ctx context.Context, input []*model.UserInput) ([]*model.User, error) {
	rets := make([]*model.User, 0)
	for _, p := range input {
		u := model.User{ID: strconv.Itoa(len(r.users) + 1), Name: p.Name}
		r.users = append(r.users, u)
		rets = append(rets, &u)
	}

	fmt.Println("create users: ", r.users)

	return rets, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Article(ctx context.Context) (*model.Article, error) {
	return &model.Article{
		ID:   "1",
		Text: "I am codinghuang",
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	rets := make([]*model.User, 0)
	for i, p := range r.users {
		rets = append(rets, &r.users[i])
		fmt.Printf("Users: %p %p\n", &r.users[i], &p)
	}
	return rets, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	for _, p := range r.users {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
