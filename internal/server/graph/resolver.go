package graph

import (
	"github.com/duckpie/cherry"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
	"google.golang.org/grpc"
)

type Connector func() (*grpc.ClientConn, error)

type Resolver struct {
	conn map[cherry.ConnKey]Connector
}

type ResolverI interface {
	Config() generated.Config
	PrepareConn(key cherry.ConnKey, conn Connector)
}

func (r *Resolver) PrepareConn(key cherry.ConnKey, conn Connector) {
	r.conn[key] = conn
}

func (r *Resolver) Config() generated.Config {
	return generated.Config{Resolvers: r}
}

func CreateResolver() *Resolver {
	return &Resolver{
		conn: make(map[cherry.ConnKey]Connector),
	}
}
