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

func (r *mutationResolver) CreateArticles(ctx context.Context, input []*model.ArticleInput) ([]*model.Article, error) {
	rets := make([]*model.Article, 0)
	for _, p := range input {
		u := model.Article{ID: strconv.Itoa(len(r.articles) + 1), Name: p.Name}
		r.articles = append(r.articles, u)
		rets = append(rets, &u)
	}

	fmt.Println("create users: ", r.articles)

	return rets, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*model.Article, error) {
	rets := make([]*model.Article, 0)
	for i, p := range r.articles {
		rets = append(rets, &r.articles[i])
		fmt.Printf("Users: %p %p\n", &r.articles[i], &p)
	}
	return rets, nil
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	for _, p := range r.articles {
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
