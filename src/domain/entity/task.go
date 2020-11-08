package entity

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	TaskID    string
	TaskName  string
	Priority  string
	LimitDate *time.Time
}

func NewTask(taskName string) *Task {
	return &Task{
		TaskID:    uuid.New().String(),
		TaskName:  taskName,
		Priority:  "Middle",
		LimitDate: nil,
	}
}

type TaskList []*Task
