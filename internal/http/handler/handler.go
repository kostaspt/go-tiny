package handler

import (
	"github.com/google/wire"

	"github.com/kostaspt/go-tiny/config"
)

var ProviderSet = wire.NewSet(NewHandler)

type Handler struct {
	config *config.Config
}

func NewHandler(c *config.Config) *Handler {
	return &Handler{
		config: c,
	}
}
