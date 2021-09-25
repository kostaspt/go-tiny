package server

import "github.com/kostaspt/go-tiny/config"

type Middleware struct {
	config *config.Config
}

func NewMiddleware(c *config.Config) *Middleware {
	return &Middleware{
		config: c,
	}
}
