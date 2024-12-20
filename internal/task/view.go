package task

type taskStatus = string

const (
	taskStatusPending taskStatus = "pending"
	taskStatusRunning taskStatus = "running"
)

type task struct {
	TaskId string     `json:"task_id"`
	Status taskStatus `json:"status"`
}
