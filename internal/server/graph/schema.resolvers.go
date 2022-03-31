package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/model"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	resp, err := r.Resolver.conn.userMs.CreateUser(ctx, &pb.NewUserReq{
		Login:    input.Login,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		UUID:      resp.Uuid,
		Login:     resp.Login,
		Email:     resp.Email,
		Role:      int(resp.Role),
		CreatedAt: resp.CreatedAt,
		UpdatedAt: resp.UpdatedAt,
	}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	dummyUser := model.User{
		UUID:      "uuid",
		Login:     "login",
		Email:     "kiwi@mail.ru",
		Role:      1,
		CreatedAt: "date",
		UpdatedAt: "date",
	}

	users = append(users, &dummyUser)
	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
