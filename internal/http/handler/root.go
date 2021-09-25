package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h Handler) RootIndex(ctx echo.Context) error {
	resp := map[string]interface{}{
		"message":  "My App API",
		"datetime": time.Now().Format(time.RFC3339Nano),
	}

	return ctx.JSON(http.StatusOK, resp)
}
