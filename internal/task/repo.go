package task

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
)

type TaskRepo interface {
	GetTask(ctx context.Context, taskId uuid.UUID) (*models.Task, error)
}

type taskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) TaskRepo {
	return &taskRepo{
		db: db,
	}
}

func (t *taskRepo) GetTask(ctx context.Context, taskId uuid.UUID) (*models.Task, error) {
	return models.Tasks(models.TaskWhere.TaskID.EQ(taskId)).One(ctx, t.db)
}
