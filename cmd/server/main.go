package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/pflag"
)

var (
	Name    = "my-app"
	Version string
)

func newApp(srv *http.Server) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger.NewKratosWrapper()),
		kratos.Server(srv),
	)
}

func main() {
	port := pflag.Uint16P("port", "p", 4000, "Server's port")
	logLevel := pflag.Int8P("log-level", "l", int8(zerolog.InfoLevel), "Logger's level of reporting")
	pflag.Parse()

	logger.Setup(Name, Version, zerolog.Level(*logLevel))

	c, err := config.New(*port)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app, cleanup, err := initApp(ctx, c)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err = app.Run(); err != nil {
		panic(err)
	}
}
