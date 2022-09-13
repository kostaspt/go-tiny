package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"

	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/scheduler"
	"github.com/kostaspt/go-tiny/pkg/logger"
)

var (
	Name    = "app-scheduler"
	Version string
)

func newApp(s *scheduler.Scheduler) (*kratos.App, func(), error) {
	if err := s.Start(); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		s.Stop()
	}

	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger.NewForKratos()),
	), cleanup, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	logLevel := pflag.Int8P("log-level", "l", int8(zerolog.InfoLevel), "Logger's level of reporting")
	pflag.Parse()

	logger.Setup(Name, Version, zerolog.Level(*logLevel))

	c, err := config.New(nil)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, cleanup, err := initApp(ctx, c)
	if err != nil {
		return err
	}
	defer cleanup()

	return app.Run()
}
