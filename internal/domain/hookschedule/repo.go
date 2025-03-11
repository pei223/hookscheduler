package hookschedule

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type HookScheduleRepo struct {
}

func NewHookScheduleRepo() *HookScheduleRepo {
	return &HookScheduleRepo{}
}

func (r *HookScheduleRepo) CreateHookSchedule(ctx context.Context, tx *sql.Tx, params *HookScheduleCreateParams) (*models.HookSchedule, error) {
	exists, err := models.Hooks(models.HookWhere.HookID.EQ(params.HookID)).Exists(ctx, tx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errorcommon.NewCommonError(errors.New("hook not found"), "hook not found", errorcommon.ErrParamsInvalid, nil)
	}
	schedule := &models.HookSchedule{
		HookScheduleID:        uuid.New(),
		HookID:                params.HookID,
		DisplayName:           params.DisplayName,
		Description:           params.Description,
		ScheduleFrequencyUnit: string(params.ScheduleFrequencyUnit),
		ScheduleTimeMonth:     int16(params.ScheduleTimeMonth),
		ScheduleTimeDay:       int16(params.ScheduleTimeDay),
		ScheduleTimeHour:      int16(params.ScheduleTimeHour),
		ScheduleTimeMinute:    int16(params.ScheduleTimeMinute),
		ScheduleTimeSecond:    int16(params.ScheduleTimeSecond),
	}
	err = schedule.Insert(ctx, tx, boil.Infer())
	return schedule, err
}

func (r *HookScheduleRepo) GetHookSchedule(ctx context.Context, tx *sql.Tx, scheduleID uuid.UUID, loadHook bool) (*models.HookSchedule, error) {
	queries := []qm.QueryMod{
		models.HookScheduleWhere.HookScheduleID.EQ(scheduleID),
	}
	if loadHook {
		queries = append(queries, qm.Load(models.HookScheduleRels.Hook))
	}
	return models.HookSchedules(queries...).One(ctx, tx)
}

func (r *HookScheduleRepo) DeleteHookSchedule(ctx context.Context, tx *sql.Tx, scheduleID uuid.UUID) error {
	schedule, err := r.GetHookSchedule(ctx, tx, scheduleID, false)
	if err != nil {
		return err
	}
	_, err = schedule.Delete(ctx, tx)
	return err
}

func (r *HookScheduleRepo) ListHookSchedules(ctx context.Context, tx *sql.Tx, listParams *types.ListParams, loadHook bool) (models.HookScheduleSlice, int, error) {
	total, err := models.HookSchedules().Count(ctx, tx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total hook schedules: %w", err)
	}
	queries := []qm.QueryMod{
		qm.Limit(listParams.Limit),
		qm.Offset(listParams.Offset),
	}
	if listParams.Sort != nil {
		queries = append(queries, qm.OrderBy(listParams.Sort.SortBy, listParams.Sort.SortOrder))
	}
	if loadHook {
		queries = append(queries, qm.Load(models.HookScheduleRels.Hook))
	}
	schedules, err := models.HookSchedules(queries...).All(ctx, tx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all hook schedules: %w", err)
	}
	return schedules, int(total), nil
}
