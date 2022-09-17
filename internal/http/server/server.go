package server

import (
	"fmt"
	"time"

	transphttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/http/handler"
	"github.com/kostaspt/go-tiny/pkg/validator"
)

var ProviderSet = wire.NewSet(NewServer, NewMiddleware)

func NewServer(c *config.Config, h *handler.Handler, m *Middleware) *transphttp.Server {
	opts := []transphttp.ServerOption{
		transphttp.Timeout(10 * time.Second),
	}

	if c.Server.Port > 0 {
		opts = append(opts, transphttp.Address(fmt.Sprintf("[::]:%d", c.Server.Port)))
	}

	e := NewEchoServer(c, h, m)

	// Prometheus
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	srv := transphttp.NewServer(opts...)

	srv.HandlePrefix("/", e)

	return srv
}

func NewEchoServer(c *config.Config, h *handler.Handler, m *Middleware) *echo.Echo {
	e := echo.New()
	e.Validator = validator.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://" + c.Domain,
			"https://" + c.Domain,
		},
	}))

	registerRoutes(e, h, m)

	return e
}

func registerRoutes(e *echo.Echo, h *handler.Handler, m *Middleware) {
	e.GET("/", h.RootIndex)
}
