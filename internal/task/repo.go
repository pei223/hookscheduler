package task

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/db"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TaskRepo interface {
	GetTask(ctx context.Context, taskId uuid.UUID) (*models.Task, error)
	CreateTask(ctx context.Context, params *TaskCreateParams) (*models.Task, error)
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

func (t *taskRepo) CreateTask(ctx context.Context, params *TaskCreateParams) (*models.Task, error) {
	var task models.Task
	err := db.ExecTx(ctx, t.db, func(ctx context.Context, tx *sql.Tx) error {
		taskID := uuid.New()
		task = models.Task{
			TaskID:   taskID,
			TaskName: params.Name,
		}
		return task.Insert(ctx, tx, boil.Infer())
	})
	return &task, err
}
