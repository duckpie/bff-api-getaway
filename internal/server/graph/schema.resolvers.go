package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wrs-news/bff-api-getaway/internal/core"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/model"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	conn := pb.NewUserServiceClient(r.Resolver.conn[core.UMS])
	resp, err := conn.CreateUser(ctx, &pb.NewUserReq{
		Login:    input.Login,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	conn := pb.NewUserServiceClient(r.Resolver.conn[core.UMS])
	resp, err := conn.UpdateUser(ctx, &pb.UpdateUserReq{
		Uuid:  input.UUID,
		Login: input.Login,
		Email: input.Email,
		Role:  int32(input.Role),
	})
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, uuid string) (*model.User, error) {
	conn := pb.NewUserServiceClient(r.Resolver.conn[core.UMS])
	resp, err := conn.DeleteUser(ctx, &pb.UserReqUuid{Uuid: uuid})
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp), nil
}

func (r *queryResolver) GetUserByUUID(ctx context.Context, uuid string) (*model.User, error) {
	conn := pb.NewUserServiceClient(r.Resolver.conn[core.UMS])
	resp, err := conn.GetUserByUuid(ctx, &pb.UserReqUuid{Uuid: uuid})
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp), nil
}

func (r *queryResolver) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	conn := pb.NewUserServiceClient(r.Resolver.conn[core.UMS])
	resp, err := conn.GetUserByLogin(ctx, &pb.UserReqLogin{Login: login})
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp), nil
}

func (r *queryResolver) GetUsersSlice(ctx context.Context, limit int, offset int) (*model.UserSelection, error) {
	conn := pb.NewUserServiceClient(r.Resolver.conn[core.UMS])
	resp, err := conn.GetAll(ctx, &pb.SelectionReq{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	return &model.UserSelection{
		Limit:    int(resp.Limit),
		Offset:   int(resp.Offset),
		Total:    int(resp.Total),
		LastPage: int(resp.LastPage),
		Data:     arrPbUserToArrGraphQlUser(resp.Data),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
