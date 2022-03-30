package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/wrs-news/bff-api-getaway/internal/config"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph"
	"github.com/wrs-news/bff-api-getaway/internal/server/graph/generated"
)

type server struct {
	cfg *config.ServerConfig
}

type ServerI interface {
	Run() error
}

func (s *server) Run() error {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", s.cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), nil))

	return nil
}

func InitServer(cfg *config.ServerConfig) ServerI {
	return &server{
		cfg: cfg,
	}
}
