package server

import (
	"fmt"

	"clean-architecture-go/infrastructure/waf/echo/middleware"
	"clean-architecture-go/infrastructure/waf/echo/router"
	"github.com/labstack/echo/v4"
)

const (
	DefaultPort = ":1323"
	BasePath    = "api"
)

type server struct {
	*echo.Echo
}

func NewServer() *server {
	return &server{
		echo.New(),
	}
}

func (server *server) Run() {
	server.routing()
	server.setMiddleware()
	server.setValidator()
	server.setErrorHandler()
	server.Logger.Fatal(server.Start(DefaultPort))
}

func (server *server) routing() {
	api := router.GetAPIs()
	for _, version := range api.Versions {
		for _, resource := range version.Resources {
			for _, endpoint := range resource.Endpoints {
				suffix := ""
				if endpoint.SuffixPath != "" {
					suffix += fmt.Sprintf("/%s", endpoint.SuffixPath)
				}
				path := fmt.Sprintf("/%s/%s/%s%s", BasePath, version.Version, resource.Resource, suffix)
				server.Router().Add(endpoint.Method, path, endpoint.HandlerFunc)
			}
		}
	}
}

func (server *server) setMiddleware() {
	for _, m := range middleware.List() {
		server.Use(m.MiddlewareFunc)
	}
}

func (server *server) setValidator() {
	//server.Validator = validator.New()
}

func (server *server) setErrorHandler() {
	//server.HTTPErrorHandler = handler.ErrorHandler
}
