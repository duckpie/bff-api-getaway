package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/duckpie/cherry"
	"github.com/wrs-news/bff-api-getaway/internal/config"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
	"google.golang.org/grpc"
)

type server struct {
	cfg *config.Config
	rlv *graph.Resolver
}

type ServerI interface {
	Run() error
	Config() (err error)
	Resolver() *graph.Resolver
}

func (s *server) Run() (err error) {
	if err := s.Config(); err != nil {
		return err
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(s.rlv.Config()))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", s.cfg.Services.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Services.Server.Port), nil))

	return nil
}

func (s *server) Config() (err error) {
	// Подключение микросевиса пользователей
	s.rlv.PrepareConn(cherry.UMS, func() (*grpc.ClientConn, error) {
		return grpc.Dial(
			fmt.Sprintf("%s:%d", s.cfg.Microservices.UserMs.Host, s.cfg.Microservices.UserMs.Port),
			grpc.WithInsecure(),
		)
	})

	// Подключение микросервиса безопасности
	s.rlv.PrepareConn(cherry.SMS, func() (*grpc.ClientConn, error) {
		return grpc.Dial(
			fmt.Sprintf("%s:%d", s.cfg.Microservices.SecurityMs.Host, s.cfg.Microservices.SecurityMs.Port),
			grpc.WithInsecure(),
		)
	})

	return
}

func (s *server) Resolver() *graph.Resolver {
	return s.rlv
}

func InitServer(cfg *config.Config) ServerI {
	return &server{
		cfg: cfg,
		rlv: graph.CreateResolver(),
	}
}
