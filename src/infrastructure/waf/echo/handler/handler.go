package handler

import (
	"clean-architecture-go/lib/actx"
	"context"
	"github.com/labstack/echo/v4"
)

func withContext(echo echo.Context, fn func(ctx context.Context) error) error {
	ctx := actx.Set(echo.Request().Context(), actx.EchoContext, echo)
	req := echo.Request().WithContext(ctx)
	echo.SetRequest(req)
	return fn(echo.Request().Context())
}
