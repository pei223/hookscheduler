package task

import (
	"github.com/gin-gonic/gin"
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
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
		return 0, nil, err
	}
	return 200, fromModel(task), nil
}

func (t *TaskWebHandler) CreateTask(c *gin.Context) (int, any, error) {
	t.logger.Info().Msg("CreateTask")
	ctx := c.Request.Context()
	params := &TaskCreateParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		t.logger.Info().Err(err).Msg("failed to bind json")
		return 400, nil, errorcommon.NewParseError(err)
	}
	invalidParams := params.Validate()
	if invalidParams != nil {
		return 0, nil, errorcommon.NewInvalidParamsError(invalidParams)
	}
	task, err := t.taskMod.CreateTask(ctx, params)
	if err != nil {
		t.logger.Error().Err(err).Msg("failed to create task")
		return 0, nil, err
	}
	return 201, fromModel(task), nil
}
