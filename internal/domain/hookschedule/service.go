package hookschedule

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/db"
	"github.com/pei223/hook-scheduler/pkg/types"
)

type HookScheduleRepoIF interface {
	GetHookSchedule(ctx context.Context, tx *sql.Tx, scheduleId uuid.UUID, loadHook bool) (*models.HookSchedule, error)
	DeleteHookSchedule(ctx context.Context, tx *sql.Tx, scheduleId uuid.UUID) error
	CreateHookSchedule(ctx context.Context, tx *sql.Tx, params *HookScheduleCreateParams) (*models.HookSchedule, error)
	ListHookSchedules(ctx context.Context, tx *sql.Tx, listParams *types.ListParams, loadHook bool) (models.HookScheduleSlice, int, error)
}

type HookScheduleService struct {
	db   *sql.DB
	repo HookScheduleRepoIF
}

func NewHookScheduleService(db *sql.DB, repo HookScheduleRepoIF) *HookScheduleService {
	return &HookScheduleService{
		repo: repo,
		db:   db,
	}
}

func (s *HookScheduleService) GetHookSchedule(ctx context.Context, scheduleId uuid.UUID, loadHook bool) (*models.HookSchedule, error) {
	var schedule *models.HookSchedule
	err := db.ReadOnlyTx(
		ctx,
		s.db,
		func(ctx context.Context, tx *sql.Tx) error {
			var err error
			schedule, err = s.repo.GetHookSchedule(ctx, tx, scheduleId, loadHook)
			return err
		},
		nil,
	)
	return schedule, err
}

func (s *HookScheduleService) DeleteHookSchedule(ctx context.Context, scheduleId uuid.UUID) error {
	err := db.ExecTx(
		ctx,
		s.db,
		func(ctx context.Context, tx *sql.Tx) error {
			return s.repo.DeleteHookSchedule(ctx, tx, scheduleId)
		},
		nil,
	)
	return err
}

func (s *HookScheduleService) CreateHookSchedule(ctx context.Context, params *HookScheduleCreateParams) (*models.HookSchedule, error) {
	var schedule *models.HookSchedule
	err := db.ExecTx(
		ctx,
		s.db,
		func(ctx context.Context, tx *sql.Tx) error {
			var err error
			schedule, err = s.repo.CreateHookSchedule(ctx, tx, params)
			return err
		},
		nil,
	)
	return schedule, err
}

func (s *HookScheduleService) ListHookSchedules(ctx context.Context, listParams *types.ListParams, loadHook bool) (models.HookScheduleSlice, int, error) {
	var schedules models.HookScheduleSlice
	var total int
	err := db.ReadOnlyTx(
		ctx,
		s.db,
		func(ctx context.Context, tx *sql.Tx) error {
			var err error
			schedules, total, err = s.repo.ListHookSchedules(ctx, tx, listParams, loadHook)
			return err
		},
		nil,
	)
	return schedules, total, err
}
