//go:build wireinject

package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/http/handler"
	"github.com/kostaspt/go-tiny/internal/http/server"
)

func initApp(context.Context, *config.Config) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			// HTTP
			server.ProviderSet,
			handler.ProviderSet,

			// Storage
			// sql.ProviderSet,
			// cache.ProviderSet,
			// wire.Bind(new(cache.Cache), new(*cache.InMemory)),
			// wire.Bind(new(cache.Cache), new(*cache.Redis)),

			newApp,
		),
	)
}
