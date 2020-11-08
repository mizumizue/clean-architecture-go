package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Middleware struct {
	MiddlewareFunc echo.MiddlewareFunc
}

func List() []*Middleware {
	return []*Middleware{
		{MiddlewareFunc: middleware.Logger()},
		{MiddlewareFunc: middleware.Recover()},
	}
}
