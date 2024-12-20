package task

import (
	"net/http"

	"github.com/rs/zerolog"
)

type TaskWebHandler struct {
	logger *zerolog.Logger
}

func NewTaskWebHandler(logger *zerolog.Logger) *TaskWebHandler {
	return &TaskWebHandler{
		logger: logger,
	}
}

func (t *TaskWebHandler) GetTask(r *http.Request) (int, any, error) {
	t.logger.Info().Msg("GetTask")
	taskId := TaskIDFromContext(r.Context())
	return 200, task{
		TaskId: taskId.String(),
		Status: taskStatusPending,
	}, nil
}
