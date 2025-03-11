package hookschedule_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/pei223/hook-scheduler/pkg/common"
	"github.com/pei223/hook-scheduler/pkg/db"
	"github.com/pei223/hook-scheduler/pkg/errorcommon"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type hookScheduleRepoTestSuite struct {
	suite.Suite
	repo   *hookschedule.HookScheduleRepo
	db     *sql.DB
	logger zerolog.Logger
}

func (s *hookScheduleRepoTestSuite) SetupSuite() {
	db := lo.Must(sql.Open("postgres", test_common.TestDatabaseConnectionString))
	s.logger = common.NewLogger(context.Background(), "debug")
	s.db = db
	s.repo = hookschedule.NewHookScheduleRepo()
}

func (s *hookScheduleRepoTestSuite) SetupTest() {
	ctx := context.TODO()
	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := models.HookSchedules().DeleteAll(ctx, s.db)
		s.Require().NoError(err)
		_, err = models.Hooks().DeleteAll(ctx, s.db)
		return err
	}, nil)
	if err != nil {
		panic(err)
	}
}

func (s *hookScheduleRepoTestSuite) TearDownSuite() {
	s.db.Close()
}

func TestHookScheduleRepo(t *testing.T) {
	suite.Run(t, new(hookScheduleRepoTestSuite))
}

func (s *hookScheduleRepoTestSuite) TestCreateHookSchedule() {
	ctx := context.TODO()

	hookID := uuid.New()
	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		model := &models.Hook{
			HookID: hookID,
			URL:    "https://example.com",
			Method: "GET",
			Body:   types.JSONB{},
			Headers: types.JSONB{
				"Content-Type": "application/json",
			},
		}
		return model.Insert(ctx, s.db, boil.Infer())
	}, nil)
	s.Require().NoError(err)

	s.Run("success", func() {
		var model *models.HookSchedule
		err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			model, err = s.repo.CreateHookSchedule(ctx, tx, &hookschedule.HookScheduleCreateParams{
				HookID:                hookID,
				DisplayName:           "test",
				Description:           "test",
				ScheduleFrequencyUnit: "every_minute",
				ScheduleTimeMinute:    0,
				ScheduleTimeSecond:    0,
				ScheduleTimeHour:      0,
				ScheduleTimeDay:       0,
				ScheduleTimeMonth:     0,
			})
			return err
		}, nil)
		s.Assert().NoError(err)
		s.Assert().NotNil(model)

		hookSchedules := lo.Must(models.HookSchedules().All(ctx, s.db))
		s.Assert().Len(hookSchedules, 1)
		s.Assert().Equal(model, hookSchedules[0])
	})

	s.Run("hook not found", func() {
		err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			_, err := s.repo.CreateHookSchedule(ctx, tx, &hookschedule.HookScheduleCreateParams{
				HookID: uuid.New(),
			})
			return err
		}, nil)
		var expectedErr *errorcommon.CommonError
		s.Require().Error(err)
		s.Assert().ErrorAs(err, &expectedErr)
	})
}

func (s *hookScheduleRepoTestSuite) TestGetHookSchedule() {
	ctx := context.TODO()

	hookID := uuid.New()
	scheduleID := uuid.New()
	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		hook := &models.Hook{
			HookID: hookID,
			URL:    "https://example.com",
			Method: "GET",
			Body:   types.JSONB{},
			Headers: types.JSONB{
				"Content-Type": "application/json",
			},
		}
		err := hook.Insert(ctx, tx, boil.Infer())
		s.Require().NoError(err)

		model := &models.HookSchedule{
			HookID:                hookID,
			HookScheduleID:        scheduleID,
			DisplayName:           "gettest",
			Description:           "gettest",
			ScheduleFrequencyUnit: "every_minute",
			ScheduleTimeMinute:    1,
			ScheduleTimeSecond:    2,
			ScheduleTimeHour:      3,
			ScheduleTimeDay:       4,
			ScheduleTimeMonth:     5,
		}
		err = model.Insert(ctx, tx, boil.Infer())
		s.Require().NoError(err)
		return nil
	}, nil)
	s.Require().NoError(err)

	s.Run("success", func() {
		var model *models.HookSchedule
		err := db.ReadOnlyTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			model, err = s.repo.GetHookSchedule(ctx, tx, scheduleID, false)
			return err
		}, nil)
		s.Assert().NoError(err)
		s.Assert().NotNil(model)
	})

	s.Run("load hook", func() {
		var model *models.HookSchedule
		err := db.ReadOnlyTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			model, err = s.repo.GetHookSchedule(ctx, tx, scheduleID, true)
			return err
		}, nil)
		s.Assert().NoError(err)
		s.Assert().NotNil(model)
		s.Assert().NotNil(model.R.Hook)
		s.Assert().Equal(hookID, model.R.Hook.HookID)
	})

	s.Run("not found", func() {
		var model *models.HookSchedule
		err := db.ReadOnlyTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			_, err := s.repo.GetHookSchedule(ctx, tx, uuid.New(), false)
			return err
		}, nil)
		s.Require().Error(err)
		s.Assert().ErrorIs(err, sql.ErrNoRows)
		s.Assert().Nil(model)
	})
}

func (s *hookScheduleRepoTestSuite) TestDeleteHookSchedule() {
	ctx := context.TODO()

	scheduleID := uuid.New()
	hookID := uuid.New()
	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		hook := &models.Hook{
			HookID: hookID,
			URL:    "https://example.com",
			Method: "GET",
			Body:   types.JSONB{},
			Headers: types.JSONB{
				"Content-Type": "application/json",
			},
		}
		err := hook.Insert(ctx, tx, boil.Infer())
		s.Require().NoError(err)

		model := &models.HookSchedule{
			HookID:                hookID,
			HookScheduleID:        scheduleID,
			DisplayName:           "deletetest",
			Description:           "deletetest",
			ScheduleFrequencyUnit: "every_minute",
			ScheduleTimeMinute:    3,
			ScheduleTimeSecond:    3,
			ScheduleTimeHour:      3,
			ScheduleTimeDay:       3,
			ScheduleTimeMonth:     3,
		}
		err = model.Insert(ctx, tx, boil.Infer())
		s.Require().NoError(err)
		return nil
	}, nil)
	s.Require().NoError(err)

	s.Run("success", func() {
		err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			return s.repo.DeleteHookSchedule(ctx, tx, scheduleID)
		}, nil)
		s.Assert().NoError(err)
		schedules := lo.Must(models.HookSchedules().All(ctx, s.db))
		s.Assert().Len(schedules, 0)
	})

	s.Run("not found", func() {
		err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			return s.repo.DeleteHookSchedule(ctx, tx, uuid.New())
		}, nil)
		s.Assert().Error(err)
		s.Assert().ErrorIs(err, sql.ErrNoRows)
	})
}

func (s *hookScheduleRepoTestSuite) TestListHookSchedules() {
	ctx := context.TODO()

	hookID := uuid.New()
	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		hook := &models.Hook{
			HookID: hookID,
			URL:    "https://example.com",
			Method: "GET",
			Body:   types.JSONB{},
			Headers: types.JSONB{
				"Content-Type": "application/json",
			},
		}
		err := hook.Insert(ctx, tx, boil.Infer())
		s.Require().NoError(err)

		for i := 0; i < 5; i++ {
			model := &models.HookSchedule{
				HookID:                hookID,
				HookScheduleID:        uuid.New(),
				DisplayName:           fmt.Sprintf("listtest%d", i),
				Description:           fmt.Sprintf("listtest%d", i),
				ScheduleFrequencyUnit: "every_minute",
				ScheduleTimeMinute:    1,
				ScheduleTimeSecond:    2,
				ScheduleTimeHour:      3,
				ScheduleTimeDay:       4,
				ScheduleTimeMonth:     5,
			}
			err = model.Insert(ctx, tx, boil.Infer())
			s.Require().NoError(err)
		}
		return nil
	}, nil)
	s.Require().NoError(err)

	s.Run("success", func() {
		var models models.HookScheduleSlice
		var total int
		err := db.ReadOnlyTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			models, total, err = s.repo.ListHookSchedules(ctx, tx, &types.ListParams{
				Limit:  3,
				Offset: 0,
			}, false)
			return err
		}, nil)
		s.Require().NoError(err)
		s.Require().Len(models, 3)
		s.Require().Equal(total, 5)
	})

	s.Run("load hook", func() {
		var models models.HookScheduleSlice
		var total int
		err := db.ReadOnlyTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			models, total, err = s.repo.ListHookSchedules(ctx, tx, &types.ListParams{
				Limit:  3,
				Offset: 0,
			}, true)
			return err
		}, nil)
		s.Require().NoError(err)
		s.Require().Len(models, 3)
		s.Require().Equal(total, 5)
		for _, model := range models {
			s.Require().NotNil(model.R.Hook)
			s.Require().Equal(hookID, model.R.Hook.HookID)
		}
	})
}
