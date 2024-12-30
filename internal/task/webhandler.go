package task

import (
	"net/http"

	"github.com/rs/zerolog"
)

type TaskWebHandler struct {
	logger  *zerolog.Logger
	taskMod TaskMod
}

func NewTaskWebHandler(logger *zerolog.Logger, taskMod TaskMod) *TaskWebHandler {
	return &TaskWebHandler{
		logger:  logger,
		taskMod: taskMod,
	}
}

func (t *TaskWebHandler) GetTask(r *http.Request) (int, any, error) {
	t.logger.Info().Msg("GetTask")
	taskId := TaskIDFromContext(r.Context())
	task, err := t.taskMod.GetTask(r.Context(), taskId)
	if err != nil {
		t.logger.Error().Err(err).Msg("failed to get task")
		return 500, nil, err
	}
	return 200, task, err
}
