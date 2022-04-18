package graph

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
	pbs "github.com/wrs-news/golang-proto/pkg/proto/security"
	pbu "github.com/wrs-news/golang-proto/pkg/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	if err := input.Validation(); err != nil {
		return nil, err
	}

	conn, err := r.Resolver.conn[cherry.UMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbu.NewUserServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := client.CreateUser(
			ctx,
			&pbu.NewUserReq{
				Login:    input.Login,
				Email:    input.Email,
				Password: input.Password,
			},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pbu.User)), nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	if err := input.Validation(); err != nil {
		return nil, err
	}

	conn, err := r.Resolver.conn[cherry.UMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	clinet := pbu.NewUserServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := clinet.UpdateUser(
			ctx,
			&pbu.UpdateUserReq{
				Uuid:  input.UUID,
				Login: input.Login,
				Email: input.Email,
				Role:  int32(input.Role),
			},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pbu.User)), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, uuid string) (*model.User, error) {
	if err := validation.Validate(uuid, validation.Required, is.UUIDv4); err != nil {
		return nil, err
	}

	conn, err := r.Resolver.conn[cherry.UMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	clinet := pbu.NewUserServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := clinet.DeleteUser(
			ctx,
			&pbu.UserReqUuid{Uuid: uuid},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pbu.User)), nil
}

func (r *mutationResolver) CreateAuth(ctx context.Context, input model.Login) (*model.Tokens, error) {
	if err := input.Validation(); err != nil {
		return nil, err
	}

	conn, err := r.Resolver.conn[cherry.SMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbs.NewSecurityServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := client.Login(
			ctx,
			&pbs.LoginReq{
				Login:    input.Login,
				Password: input.Password,
			},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Tokens{
		RefreshToken: resp.(*pbs.TokensPair).RefreshToken,
		AccessToken:  resp.(*pbs.TokensPair).AccessToken,
	}, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, token string) (*model.Tokens, error) {
	conn, err := r.Resolver.conn[cherry.SMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbs.NewSecurityServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := client.RefreshToken(
			ctx,
			&pbs.RefreshTokenReq{Token: token},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return &model.Tokens{
		RefreshToken: resp.(*pbs.TokensPair).RefreshToken,
		AccessToken:  resp.(*pbs.TokensPair).AccessToken,
	}, nil
}

func (r *mutationResolver) Logout(ctx context.Context, accessToken string) (*model.StatusResp, error) {
	conn, err := r.Resolver.conn[cherry.SMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbs.NewSecurityServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := client.Logout(
			ctx,
			&pbs.LogoutReq{Token: accessToken},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return &model.StatusResp{Status: string(cherry.StatusFail)}, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	if _, err := rptr(ctx); err != nil {
		return &model.StatusResp{Status: string(cherry.StatusFail)}, err
	}

	return &model.StatusResp{Status: string(cherry.StatusOk)}, nil
}

func (r *queryResolver) GetUserByUUID(ctx context.Context, uuid string) (*model.User, error) {
	if err := validation.Validate(uuid, validation.Required, is.UUIDv4); err != nil {
		return nil, err
	}

	conn, err := r.Resolver.conn[cherry.UMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbu.NewUserServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := client.GetUserByUuid(
			ctx,
			&pbu.UserReqUuid{Uuid: uuid},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pbu.User)), nil
}

func (r *queryResolver) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	if err := validation.Validate(login, validation.Required,
		validation.Match(regexp.MustCompile(cherry.RegexName))); err != nil {
		return nil, err
	}

	conn, err := r.Resolver.conn[cherry.UMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbu.NewUserServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := client.GetUserByLogin(
			ctx,
			&pbu.UserReqLogin{Login: login},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return pbUserToGraphQlUser(resp.(*pbu.User)), nil
}

func (r *queryResolver) GetUsersSlice(ctx context.Context, limit int, offset int) (*model.UserSelection, error) {
	if err := validation.Validate(limit, validation.Min(1),
		validation.Max(30), validation.Required); err != nil {
		return nil, err
	}

	if err := validation.Validate(offset, validation.Min(0),
		validation.Required.When(offset > 0)); err != nil {
		return nil, err
	}

	conn, err := r.Resolver.conn[cherry.UMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbu.NewUserServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := client.GetAll(
			ctx,
			&pbu.SelectionReq{
				Limit:  int32(limit),
				Offset: int32(offset),
			},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	resp, err := rptr(ctx)
	if err != nil {
		return nil, err
	}

	return &model.UserSelection{
		Limit:    int(resp.(*pbu.Selection).Limit),
		Offset:   int(resp.(*pbu.Selection).Offset),
		Total:    int(resp.(*pbu.Selection).Total),
		LastPage: int(resp.(*pbu.Selection).LastPage),
		Data:     arrPbUserToArrGraphQlUser(resp.(*pbu.Selection).Data),
	}, nil
}

func (r *queryResolver) AuthCheck(ctx context.Context, accessToken string) (*model.StatusResp, error) {
	if err := validation.Validate(accessToken, validation.Required); err != nil {
		return &model.StatusResp{Status: string(cherry.StatusFail)}, err
	}

	conn, err := r.Resolver.conn[cherry.SMS]()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pbs.NewSecurityServiceClient(conn)
	rptr := cherrynet.GrpcRepeater(func(ctx context.Context) (interface{}, error) {
		resp, err := client.AuthCheck(
			ctx,
			&pbs.AuthCheckReq{AccessToken: accessToken},
			grpc.UseCompressor(gzip.Name),
		)
		if err != nil {
			return &model.StatusResp{Status: string(cherry.StatusFail)}, err
		}

		return resp, nil
	}, 2, time.Millisecond*500)

	if _, err := rptr(ctx); err != nil {
		return &model.StatusResp{Status: string(cherry.StatusFail)}, err
	}

	return &model.StatusResp{Status: string(cherry.StatusOk)}, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
