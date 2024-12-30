package task

import (
	"github.com/gin-gonic/gin"
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

func (t *TaskWebHandler) GetTask(c *gin.Context) (int, any, error) {
	t.logger.Info().Msg("GetTask")
	ctx := c.Request.Context()
	taskId := TaskIDFromContext(ctx)
	task, err := t.taskMod.GetTask(ctx, taskId)
	if err != nil {
		t.logger.Error().Err(err).Msg("failed to get task")
		return 500, nil, err
	}
	return 200, task, err
}
