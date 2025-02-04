package webapi

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hook"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
	"github.com/pei223/hook-scheduler/pkg/logger"
)

type HookUsecaseIF interface {
	GetHook(ctx context.Context, hookId uuid.UUID) (*models.Hook, error)
	DeleteHook(ctx context.Context, hookId uuid.UUID) error
	CreateHook(ctx context.Context, params *hook.HookCreateParams) (*models.Hook, error)
}

type HookRouter struct {
	hookUsecase HookUsecaseIF
}

func NewHookRouter(hookUsecase HookUsecaseIF) *HookRouter {
	return &HookRouter{
		hookUsecase: hookUsecase,
	}
}

func (t *HookRouter) GetHook(c *gin.Context) (int, any, error) {
	ctx := c.Request.Context()
	logger := logger.FromContext(ctx)
	logger.Info().Msg("get hook")
	hookId := HookIDFromContext(ctx)
	hook, err := t.hookUsecase.GetHook(ctx, hookId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.Error().Err(err).Msg("failed to get hook")
		}
		return 0, nil, err
	}
	return 200, fromModel(hook), nil
}

func (t *HookRouter) DeleteHook(c *gin.Context) (int, any, error) {
	ctx := c.Request.Context()
	logger := logger.FromContext(ctx)
	hookId := HookIDFromContext(ctx)
	err := t.hookUsecase.DeleteHook(ctx, hookId)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get hook")
		return 0, nil, err
	}
	return http.StatusNoContent, nil, nil
}

func (t *HookRouter) CreateHook(c *gin.Context) (int, any, error) {
	ctx := c.Request.Context()
	logger := logger.FromContext(ctx)
	params := &hook.HookCreateParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		logger.Info().Err(err).Msg("failed to bind json")
		return 400, nil, errorcommon.NewParseError(err)
	}
	invalidParams := params.Validate()
	if invalidParams != nil {
		return 0, nil, errorcommon.NewInvalidParamsError(invalidParams)
	}
	hook, err := t.hookUsecase.CreateHook(ctx, params)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create hook")
		return 0, nil, err
	}
	return 201, fromModel(hook), nil
}
