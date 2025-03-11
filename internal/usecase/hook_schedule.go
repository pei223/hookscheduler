package usecase

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/types"
)

type HookScheduleServiceIF interface {
	CreateHookSchedule(ctx context.Context, params *hookschedule.HookScheduleCreateParams) (*models.HookSchedule, error)
	ListHookSchedules(ctx context.Context, listParams *types.ListParams, loadHook bool) (models.HookScheduleSlice, int, error)
	GetHookSchedule(ctx context.Context, scheduleId uuid.UUID, loadHook bool) (*models.HookSchedule, error)
	DeleteHookSchedule(ctx context.Context, scheduleId uuid.UUID) error
}

type HookScheduleUsecase struct {
	db      *sql.DB
	service HookScheduleServiceIF
}

func NewHookScheduleUsecase(db *sql.DB, service HookScheduleServiceIF) *HookScheduleUsecase {
	return &HookScheduleUsecase{db: db, service: service}
}

func (u *HookScheduleUsecase) CreateHookSchedule(ctx context.Context, params *hookschedule.HookScheduleCreateParams) (*models.HookSchedule, error) {
	return u.service.CreateHookSchedule(ctx, params)
}

func (u *HookScheduleUsecase) ListHookSchedules(ctx context.Context, listParams *types.ListParams, loadHook bool) (models.HookScheduleSlice, int, error) {
	return u.service.ListHookSchedules(ctx, listParams, loadHook)
}

func (u *HookScheduleUsecase) GetHookSchedule(ctx context.Context, scheduleId uuid.UUID, loadHook bool) (*models.HookSchedule, error) {
	return u.service.GetHookSchedule(ctx, scheduleId, loadHook)
}

func (u *HookScheduleUsecase) DeleteHookSchedule(ctx context.Context, scheduleId uuid.UUID) error {
	return u.service.DeleteHookSchedule(ctx, scheduleId)
}
