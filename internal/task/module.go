package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
)

type TaskMod interface {
	GetTask(ctx context.Context, taskId uuid.UUID) (*models.Task, error)
	CreateTask(ctx context.Context, params *TaskCreateParams) (*models.Task, error)
}

type taskMod struct {
	taskRepo TaskRepo
}

func NewTaskMod(taskRepo TaskRepo) TaskMod {
	return &taskMod{
		taskRepo: taskRepo,
	}
}

func (t *taskMod) GetTask(ctx context.Context, taskId uuid.UUID) (*models.Task, error) {
	return t.taskRepo.GetTask(ctx, taskId)
}

func (t *taskMod) CreateTask(ctx context.Context, params *TaskCreateParams) (*models.Task, error) {
	return t.taskRepo.CreateTask(ctx, params)
}
