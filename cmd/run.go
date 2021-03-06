package cmd

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/wrs-news/bff-api-getaway/internal/config"
	"github.com/wrs-news/bff-api-getaway/internal/server"
	"golang.org/x/sync/errgroup"
)

func runCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run microservice",
		Long:  `...`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.NewConfig()

			if _, err := toml.DecodeFile(
				fmt.Sprintf("config/config.%s.toml", os.Getenv("ENV")), cfg); err != nil {
				log.Printf(err.Error())
				os.Exit(1)
			}

			if err := runner(cfg); err != nil {
				log.Printf(err.Error())
				os.Exit(1)
			}
		},
	}

	flag.Parse()
	return cmd
}

func runner(cfg *config.Config) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
		}
	}()

	// Создание контекста
	ctx := context.Background()
	errs, ctx := errgroup.WithContext(ctx)

	errs.Go(func() error {
		return server.InitServer(cfg).Run()
	})

	cancelInterrupt := make(chan struct{})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-c:
		return fmt.Errorf("received signal %s", sig)
	case <-cancelInterrupt:
		return errs.Wait()
	}
}
