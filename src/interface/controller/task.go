package controller

import (
	"context"

	"clean-architecture-go/application/usecase"
)

type TaskController struct {
	use usecase.ITaskUseCase
}

func NewTaskController(use usecase.ITaskUseCase) *TaskController {
	return &TaskController{use: use}
}

func (con *TaskController) ListTask(ctx context.Context) error {
	return con.use.ListHandle(ctx)
}

func (con *TaskController) CreateTask(ctx context.Context, taskName string) error {
	inputData := usecase.NewTaskCreateInputData(taskName)
	return con.use.CreateHandle(ctx, inputData)
}
