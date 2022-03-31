package graph

import (
	"fmt"

	"github.com/wrs-news/bff-api-getaway/internal/config"
	pb "github.com/wrs-news/golang-proto/pkg/proto/user"
	"google.golang.org/grpc"
)

type Connections struct {
	userMs pb.UserServiceClient
}

type Resolver struct {
	conn *Connections
}

type ResolverI interface {
	SetConnections(cfg *config.MicroservicesConfigs) error
}

func (r *Resolver) SetConnections(cfg *config.MicroservicesConfigs) (func() error, error) {
	userConn, err := grpc.Dial(fmt.Sprintf(":%d", cfg.UserMs.Port), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	r.conn = &Connections{
		userMs: pb.NewUserServiceClient(userConn),
	}

	return func() error {
		return userConn.Close()
	}, nil
}

func CreateResolver() *Resolver {
	return &Resolver{}
}
