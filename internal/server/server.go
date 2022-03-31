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
	cfg *config.Config
}

type ServerI interface {
	Run() error
}

func (s *server) Run() error {
	res := graph.CreateResolver()
	breaker, err := res.SetConnections(&s.cfg.Microservices)
	if err != nil {
		return err
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: res}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", s.cfg.Services.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Services.Server.Port), nil))

	return breaker()
}

func InitServer(cfg *config.Config) ServerI {
	return &server{
		cfg: cfg,
	}
}
