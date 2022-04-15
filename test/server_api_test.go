//go:build api_test

package api_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/assert"

	"github.com/kostaspt/go-tiny/config"
	"github.com/kostaspt/go-tiny/internal/http/handler"
	"github.com/kostaspt/go-tiny/internal/http/server"
	"github.com/kostaspt/go-tiny/internal/storage/sql"
)

func StartServer(ctx context.Context, t *testing.T) (*httpexpect.Expect, *sql.DB, func()) {
	c, err := config.New(nil)
	assert.NoError(t, err)

	s, cleanup, err := sql.NewConnection(ctx, c)
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

	return e, s, func() {
		cleanup()
		testServer.Close()
	}
}
