package input

import (
	"context"

	"clean-architecture-go/domain/entity"
)

type ITaskRepository interface {
	List(ctx context.Context) (entity.TaskList, error)
	Get(ctx context.Context, taskID string) error
	Create(ctx context.Context, task *entity.Task) error
	Update(ctx context.Context, task *entity.Task) error
	Delete(ctx context.Context, taskID string) error
}
