package firestore

import (
	"clean-architecture-go/application/input"
	"clean-architecture-go/domain/entity"
	"cloud.google.com/go/firestore"
	"context"
)

var _ input.ITaskRepository = &TaskFirestoreRepository{}

type TaskFirestoreRepository struct{}

func NewTaskFirestoreRepository() input.ITaskRepository {
	return &TaskFirestoreRepository{}
}

func (repo *TaskFirestoreRepository) List(ctx context.Context) (taskList entity.TaskList, err error) {
	_ = withInitializedClient(ctx, func(fs *firestore.Client) error {
		taskList = entity.TaskList{
			&entity.Task{
				TaskID:    "11",
				TaskName:  "hogehoge",
				Priority:  "Middle",
				LimitDate: nil,
			},
		}
		return nil
	})
	return taskList, nil
}

func (repo *TaskFirestoreRepository) Get(ctx context.Context, taskID string) error {
	panic("implement me")
}

func (repo *TaskFirestoreRepository) Create(ctx context.Context, task *entity.Task) error {
	panic("implement me")
}

func (repo *TaskFirestoreRepository) Update(ctx context.Context, task *entity.Task) error {
	panic("implement me")
}

func (repo TaskFirestoreRepository) Delete(ctx context.Context, taskID string) error {
	panic("implement me")
}
