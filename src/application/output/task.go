package output

import (
	"context"
	"time"
)

type TaskCreateOutputData struct {
	TaskID    string
	TaskName  string
	Priority  string
	LimitDate *time.Time
}

type TaskGetOutputData struct {
	TaskID    string
	TaskName  string
	Priority  string
	LimitDate *time.Time
}

func NewTaskGetOutputData(taskID string, taskName string, priority string, limitDate *time.Time) *TaskGetOutputData {
	return &TaskGetOutputData{TaskID: taskID, TaskName: taskName, Priority: priority, LimitDate: limitDate}
}

type TaskListOutputData []*TaskGetOutputData

type ITaskPresenter interface {
	List(ctx context.Context, outputData TaskListOutputData) error
	Created(ctx context.Context, outputData *TaskCreateOutputData) error
}
