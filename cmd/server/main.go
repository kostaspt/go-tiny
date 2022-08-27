package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"

	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/pkg/logger"
)

var (
	Name    = "my-app"
	Version string
)

func newApp(srv *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger.NewForKratos()),
		kratos.Server(srv),
	)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	port := pflag.Uint16P("port", "p", 4000, "Server's port")
	logLevel := pflag.Int8P("log-level", "l", int8(zerolog.InfoLevel), "Logger's level of reporting")
	pflag.Parse()

	logger.Setup(Name, Version, zerolog.Level(*logLevel))

	c, err := config.New(port)
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
