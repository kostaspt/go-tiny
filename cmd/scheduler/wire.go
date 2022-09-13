//go:build wireinject

package main

import (
	"context"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/scheduler"
	"github.com/kostaspt/go-tiny/internal/scheduler/job"
	"github.com/kostaspt/go-tiny/internal/scheduler/registrar"
)

func initApp(context.Context, *config.Config) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			job.ProviderSet,
			registrar.ProviderSet,
			scheduler.ProviderSet,

			newApp,
		),
	)
}
