package taskstore

import (
	"context"

	"github.com/task-manager-cli/cmd/taskmanager"
)

type Store interface {
	Insert(ctx context.Context, task *taskmanager.Task) error
	Retrieve(ctx context.Context, id int) (*taskmanager.Task, error)
	Update(ctx context.Context, task *taskmanager.Task) error
	Delete(ctx context.Context, id int) error
	List(ctx context.Context) ([]*taskmanager.Task, error)
	Persist(ctx context.Context) error
}
