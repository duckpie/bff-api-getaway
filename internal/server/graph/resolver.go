package graph

import (
	"github.com/duckpie/cherry"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
	"google.golang.org/grpc"
)

type Resolver struct {
	conn map[cherry.ConnKey]*grpc.ClientConn
}

type ResolverI interface {
	Config() generated.Config
	AddConnection(key cherry.ConnKey, connect func() (*grpc.ClientConn, error)) (err error)
	Clear() (err error)
}

func (r *Resolver) AddConnection(key cherry.ConnKey, connect func() (*grpc.ClientConn, error)) (err error) {
	conn, err := connect()
	if err != nil {
		return err
	}

	r.conn[key] = conn
	return
}

func (r *Resolver) Config() generated.Config {
	return generated.Config{Resolvers: r}
}

func (r *Resolver) Clear() (err error) {
	for _, c := range r.conn {
		return c.Close()
	}

	return
}

func CreateResolver() *Resolver {
	return &Resolver{
		conn: make(map[cherry.ConnKey]*grpc.ClientConn),
	}
}
