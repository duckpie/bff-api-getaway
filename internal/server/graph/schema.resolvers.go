package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"regexp"
	"time"

	"github.com/duckpie/cherry"
	cherrynet "github.com/duckpie/cherry/net"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/model"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	if err := input.Validation(); err != nil {
		return nil, err
	}

	conn := pb.NewUserServiceClient(r.Resolver.conn[cherry.UMS])
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := conn.CreateUser(ctx, &pb.NewUserReq{
			Login:    input.Login,
			Email:    input.Email,
			Password: input.Password,
		})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 3, time.Second)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pb.User)), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	if err := input.Validation(); err != nil {
		return nil, err
	}

	conn := pb.NewUserServiceClient(r.Resolver.conn[cherry.UMS])
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := conn.UpdateUser(ctx, &pb.UpdateUserReq{
			Uuid:  input.UUID,
			Login: input.Login,
			Email: input.Email,
			Role:  int32(input.Role),
		})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 3, time.Second)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pb.User)), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, uuid string) (*model.User, error) {
	if err := validation.Validate(uuid, validation.Required, is.UUIDv4); err != nil {
		return nil, err
	}

	conn := pb.NewUserServiceClient(r.Resolver.conn[cherry.UMS])
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := conn.DeleteUser(ctx, &pb.UserReqUuid{Uuid: uuid})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 3, time.Second)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pb.User)), nil
}

func (r *queryResolver) GetUserByUUID(ctx context.Context, uuid string) (*model.User, error) {
	if err := validation.Validate(uuid, validation.Required, is.UUIDv4); err != nil {
		return nil, err
	}

	conn := pb.NewUserServiceClient(r.Resolver.conn[cherry.UMS])
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := conn.GetUserByUuid(ctx, &pb.UserReqUuid{Uuid: uuid})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 3, time.Second)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pb.User)), nil
}

func (r *queryResolver) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	if err := validation.Validate(login, validation.Required, validation.Match(regexp.MustCompile(cherry.RegexName))); err != nil {
		return nil, err
	}

	conn := pb.NewUserServiceClient(r.Resolver.conn[cherry.UMS])
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := conn.GetUserByLogin(ctx, &pb.UserReqLogin{Login: login})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 3, time.Second)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pb.User)), nil
}

func (r *queryResolver) GetUsersSlice(ctx context.Context, limit int, offset int) (*model.UserSelection, error) {
	if err := validation.Validate(limit, validation.Min(1), validation.Max(30), validation.Required); err != nil {
		return nil, err
	}

	if err := validation.Validate(offset, validation.Min(0), validation.Required.When(offset > 0)); err != nil {
		return nil, err
	}

	conn := pb.NewUserServiceClient(r.Resolver.conn[cherry.UMS])
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := conn.GetAll(ctx, &pb.SelectionReq{
			Limit:  int32(limit),
			Offset: int32(offset),
		})
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 3, time.Second)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return &model.UserSelection{
		Limit:    int(resp.(*pb.Selection).Limit),
		Offset:   int(resp.(*pb.Selection).Offset),
		Total:    int(resp.(*pb.Selection).Total),
		LastPage: int(resp.(*pb.Selection).LastPage),
		Data:     arrPbUserToArrGraphQlUser(resp.(*pb.Selection).Data),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
