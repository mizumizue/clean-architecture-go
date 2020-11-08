package usecase

import (
	"context"

	"clean-architecture-go/application/input"
	"clean-architecture-go/application/output"
)

type ITaskUseCase interface {
	ListHandle(ctx context.Context) error
	CreateHandle(ctx context.Context, inputData *TaskCreateInputData) error
}

var _ ITaskUseCase = &TaskUseCase{}

type TaskUseCase struct {
	repo input.ITaskRepository
	pre  output.ITaskPresenter
}

func NewTaskUseCase(repo input.ITaskRepository, pre output.ITaskPresenter) ITaskUseCase {
	return &TaskUseCase{repo: repo, pre: pre}
}

func (t *TaskUseCase) ListHandle(ctx context.Context) error {
	list, err := t.repo.List(ctx)
	if err != nil {
		return err
	}
	outputData := make(output.TaskListOutputData, len(list))
	for i, task := range list {
		t := output.NewTaskGetOutputData(task.TaskID, task.TaskName, task.Priority, task.LimitDate)
		outputData[i] = t
	}
	return t.pre.List(ctx, outputData)
}

func (t *TaskUseCase) CreateHandle(ctx context.Context, inputData *TaskCreateInputData) error {
	panic("implement me")
}

type TaskCreateInputData struct {
	TaskName string
}

func NewTaskCreateInputData(taskName string) *TaskCreateInputData {
	return &TaskCreateInputData{TaskName: taskName}
}
