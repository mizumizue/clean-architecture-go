package router

import (
	"clean-architecture-go/di"
	"github.com/labstack/echo/v4"
)

type API struct {
	Versions []*Version
}

type Version struct {
	Version   string
	Resources []*Resource
}

type Resource struct {
	Resource  string
	Endpoints []*Endpoint
}

type Endpoint struct {
	Method      string
	SuffixPath  string
	HandlerFunc echo.HandlerFunc
}

func GetAPIs() *API {
	// Handlers
	taskHandler := di.InitializeTaskHandler()

	return &API{
		Versions: []*Version{
			{
				Version: "v1",
				Resources: []*Resource{
					{
						Resource: "tasks",
						Endpoints: []*Endpoint{
							{Method: echo.GET, SuffixPath: "", HandlerFunc: taskHandler.List},
							{Method: echo.GET, SuffixPath: ":task_id", HandlerFunc: taskHandler.Get},
							{Method: echo.POST, SuffixPath: "", HandlerFunc: taskHandler.Create},
							{Method: echo.PUT, SuffixPath: ":task_id", HandlerFunc: taskHandler.Update},
							{Method: echo.PATCH, SuffixPath: ":task_id", HandlerFunc: taskHandler.Update},
							{Method: echo.DELETE, SuffixPath: ":task_id", HandlerFunc: taskHandler.Delete},
						},
					},
				},
			},
		},
	}
}
