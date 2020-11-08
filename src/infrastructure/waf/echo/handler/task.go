package handler

import (
	"clean-architecture-go/interface/controller"
	"context"
	"github.com/labstack/echo/v4"
)

type ITaskHandler interface {
	List(echo echo.Context) error
	Get(echo echo.Context) error
	Create(echo echo.Context) error
	Update(echo echo.Context) error
	Delete(echo echo.Context) error
}

var _ ITaskHandler = &TaskHandler{}

type TaskHandler struct {
	con *controller.TaskController
}

func NewTaskHandler(con *controller.TaskController) ITaskHandler {
	return &TaskHandler{con: con}
}

func (h *TaskHandler) List(echo echo.Context) error {
	return withContext(echo, func(ctx context.Context) error {
		return h.con.ListTask(ctx)
	})
}

func (h *TaskHandler) Get(echo echo.Context) error {
	panic("implement me")
}

func (h *TaskHandler) Create(echo echo.Context) error {
	panic("implement me")
}

func (h *TaskHandler) Update(echo echo.Context) error {
	panic("implement me")
}

func (h *TaskHandler) Delete(echo echo.Context) error {
	panic("implement me")
}
