package jsonp

import (
	"context"
	"net/http"

	"clean-architecture-go/application/output"
)

var _ output.ITaskPresenter = &TaskJsonPresenter{}

type TaskJsonPresenter struct {
	jw JsonWriter
}

func NewTaskJsonPresenter() output.ITaskPresenter {
	return &TaskJsonPresenter{}
}

func (pre *TaskJsonPresenter) List(ctx context.Context, outputData output.TaskListOutputData) error {
	return withJsonWriter(ctx, func(jw JsonWriter) error {
		return jw.JSON(http.StatusOK, outputData)
	})
}

func (pre *TaskJsonPresenter) Created(ctx context.Context, outputData *output.TaskCreateOutputData) error {
	return withJsonWriter(ctx, func(jw JsonWriter) error {
		return jw.JSON(http.StatusOK, outputData)
	})
}
