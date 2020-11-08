package memory

import (
	"context"

	"clean-architecture-go/application/input"
	"clean-architecture-go/domain/entity"
)

var _ input.ITaskRepository = &TaskInMemoryRepository{}

type TaskInMemoryRepository struct{}

func NewTaskInMemoryRepository() input.ITaskRepository {
	return &TaskInMemoryRepository{}
}

func (repo *TaskInMemoryRepository) List(ctx context.Context) (entity.TaskList, error) {
	panic("implement me")
}

func (repo *TaskInMemoryRepository) Get(ctx context.Context, taskID string) error {
	panic("implement me")
}

func (repo *TaskInMemoryRepository) Create(ctx context.Context, task *entity.Task) error {
	panic("implement me")
}

func (repo *TaskInMemoryRepository) Update(ctx context.Context, task *entity.Task) error {
	panic("implement me")
}

func (repo TaskInMemoryRepository) Delete(ctx context.Context, taskID string) error {
	panic("implement me")
}
