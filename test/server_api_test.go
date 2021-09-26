//go:build api_test

package api_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/http/handler"
	"github.com/kostaspt/go-tiny/internal/http/server"
	"github.com/kostaspt/go-tiny/internal/sql"
	"github.com/kostaspt/go-tiny/internal/sql/ent"
	"github.com/stretchr/testify/assert"
)

func StartServer(ctx context.Context, t *testing.T) (*httpexpect.Expect, *ent.Client, func(), func()) {
	c, err := config.New(0)
	assert.NoError(t, err)

	s, cleanup, err := sql.NewClient(ctx, c)
	assert.NoError(t, err)

	h := handler.NewHandler(c)
	m := server.NewMiddleware(c)
	srv := server.NewEchoServer(c, h, m)

	testServer := httptest.NewServer(srv)

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  testServer.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			// Enable for debugging only
			// httpexpect.NewDebugPrinter(t, true),
		},
	})

	return e, s,
		func() {
			s.User.Delete().Exec(ctx)
		}, func() {
			cleanup()
			testServer.Close()
		}
}
