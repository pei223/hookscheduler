package task

import "github.com/pei223/hook-scheduler/internal/models"

type taskStatus = string

const (
	taskStatusPending taskStatus = "pending"
	taskStatusRunning taskStatus = "running"
)

type Task struct {
	TaskId string     `json:"task_id"`
	Name   string     `json:"name"`
	Status taskStatus `json:"status"`
}

func fromModel(model *models.Task) Task {
	return Task{
		TaskId: model.TaskID.String(),
		Name:   model.TaskName,
		Status: taskStatusPending,
	}
}
