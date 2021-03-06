package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gcode/graph/generated"
	"gcode/graph/models"
	"strconv"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	// panic(fmt.Errorf("not implemented"))
	return userCreateUpdate(r, input, false)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input models.UserInput) (*models.User, error) {
	// panic(fmt.Errorf("not implemented"))
	return userCreateUpdate(r, input, true, id)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	// panic(fmt.Errorf("not implemented"))
	return userDelete(r, id)
}

func (r *mutationResolver) CreateArticles(ctx context.Context, input []*models.ArticleInput) ([]*models.Article, error) {
	rets := make([]*models.Article, 0)
	for _, p := range input {
		u := models.Article{ID: strconv.Itoa(len(r.articles) + 1), Name: p.Name}
		r.articles = append(r.articles, u)
		rets = append(rets, &u)
	}

	fmt.Println("create users: ", r.articles)

	return rets, nil
}

func (r *queryResolver) Auth(ctx context.Context, input models.UserAuth) (bool, error) {
	return userAuth(r, &input)
}

func (r *queryResolver) Users(ctx context.Context, id *string) (*models.Users, error) {
	// fmt.Println("1UserID: ", id)
	// return userList(r, id)
	usrid := string("123")
	usr := models.User{
		ID:     "ec17af15-e354-440c-a09f-69715fc8b595",
		Email:  "your@email.com",
		UserID: &usrid,
	}

	list := make([]*models.User, 0)
	list = append(list, &usr)
	num := 1

	records := &models.Users{Count: &num, List: list}
	return records, nil
}

func (r *queryResolver) Articles(ctx context.Context) ([]*models.Article, error) {
	fmt.Println("articles...")
	rets := make([]*models.Article, 0)
	for i, p := range r.articles {
		rets = append(rets, &r.articles[i])
		fmt.Printf("Users: %p %p\n", &r.articles[i], &p)
	}
	return rets, nil
}

func (r *queryResolver) Article(ctx context.Context, id string) (*models.Article, error) {
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
