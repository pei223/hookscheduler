package webapi

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
	"github.com/pei223/hook-scheduler/pkg/logger"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/pei223/hook-scheduler/pkg/web"
)

type HookScheduleUsecaseIF interface {
	GetHookSchedule(ctx context.Context, scheduleId uuid.UUID, loadHook bool) (*models.HookSchedule, error)
	DeleteHookSchedule(ctx context.Context, scheduleId uuid.UUID) error
	CreateHookSchedule(ctx context.Context, params *hookschedule.HookScheduleCreateParams) (*models.HookSchedule, error)
	ListHookSchedules(ctx context.Context, listParams *types.ListParams, loadHook bool) (models.HookScheduleSlice, int, error)
}

type HookScheduleRouter struct {
	hookScheduleUsecase HookScheduleUsecaseIF
}

func NewHookScheduleRouter(hookScheduleUsecase HookScheduleUsecaseIF) *HookScheduleRouter {
	return &HookScheduleRouter{hookScheduleUsecase: hookScheduleUsecase}
}

func (r *HookScheduleRouter) GetHookSchedule(c *gin.Context) (int, any, error) {
	ctx := c.Request.Context()
	logger := logger.FromContext(ctx)
	hookScheduleID := HookScheduleIDFromContext(c.Request.Context())
	hookSchedule, err := r.hookScheduleUsecase.GetHookSchedule(c.Request.Context(), hookScheduleID, false)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			logger.Error().Err(err).Msg("failed to get hook")
		}
		return 0, nil, err
	}
	return 200, fromHookScheduleModel(hookSchedule), nil
}

func (r *HookScheduleRouter) CreateHookSchedule(c *gin.Context) (int, any, error) {
	ctx := c.Request.Context()
	logger := logger.FromContext(ctx)
	params := &hookschedule.HookScheduleCreateParams{}
	if err := c.ShouldBindJSON(params); err != nil {
		logger.Info().Err(err).Msg("failed to bind json")
		return 400, nil, errorcommon.NewParseError(err)
	}
	invalidParams := params.Validate()
	if invalidParams != nil {
		return 0, nil, errorcommon.NewInvalidParamsError(invalidParams)
	}
	hookSchedule, err := r.hookScheduleUsecase.CreateHookSchedule(ctx, params)
	if err != nil {
		logger.Error().Err(err).Msg("failed to create hook schedule")
		return 0, nil, err
	}
	return 201, fromHookScheduleModel(hookSchedule), nil
}

func (r *HookScheduleRouter) DeleteHookSchedule(c *gin.Context) (int, any, error) {
	ctx := c.Request.Context()
	logger := logger.FromContext(ctx)
	hookScheduleID := HookScheduleIDFromContext(c.Request.Context())
	err := r.hookScheduleUsecase.DeleteHookSchedule(ctx, hookScheduleID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to delete hook schedule")
		return 0, nil, err
	}
	return 204, nil, nil
}

func (r *HookScheduleRouter) ListHookSchedules(c *gin.Context) (int, any, error) {
	ctx := c.Request.Context()
	logger := logger.FromContext(ctx)

	limit, invalidParam := web.ValidateInt(c, "limit", 10)
	if invalidParam != nil {
		return 0, nil, errorcommon.NewInvalidParamsError([]errorcommon.InvalidParam{*invalidParam})
	}
	offset, invalidParam := web.ValidateInt(c, "offset", 0)
	if invalidParam != nil {
		return 0, nil, errorcommon.NewInvalidParamsError([]errorcommon.InvalidParam{*invalidParam})
	}

	hookSchedules, total, err := r.hookScheduleUsecase.ListHookSchedules(ctx, &types.ListParams{Limit: limit, Offset: offset}, false)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get all hook schedules")
		return 0, nil, err
	}
	return 200, fromHookScheduleModels(hookSchedules, total, limit, offset), nil
}
