package hook_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pei223/hook-scheduler/internal/domain/hook"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/pei223/hook-scheduler/pkg/common"
	"github.com/pei223/hook-scheduler/pkg/db"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type hookRepoTestSuite struct {
	suite.Suite
	repo   hook.HookRepo
	db     *sql.DB
	logger zerolog.Logger
}

func (s *hookRepoTestSuite) SetupSuite() {
	db := lo.Must(sql.Open("postgres", test_common.TestDatabaseConnectionString))
	s.logger = common.NewLogger(context.Background(), "debug")
	s.db = db
	s.repo = hook.NewHookRepo()
}

func (s *hookRepoTestSuite) SetupTest() {
	ctx := context.TODO()
	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := models.Hooks().DeleteAll(ctx, s.db)
		return err
	}, nil)
	if err != nil {
		panic(err)
	}
}

func (s *hookRepoTestSuite) TearDownSuite() {
	s.db.Close()
}

func TestHookRepoSuite(t *testing.T) {
	suite.Run(t, new(hookRepoTestSuite))
}

func (s *hookRepoTestSuite) TestCreateHook() {
	ctx := context.TODO()

	s.Run("success", func() {
		var model *models.Hook
		err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			var err error
			model, err = s.repo.CreateHook(ctx, tx, &hook.HookCreateParams{
				URL:    "http://example.com",
				Method: "POST",
				Body: types.JSONB{
					"testKey": "testValue",
				},
				Headers: types.JSONB{
					"Authorization": "Bearer hoge",
				},
			})
			return err
		}, nil)
		s.Assert().NoError(err)
		s.Assert().NotNil(model)

		hooks := lo.Must(models.Hooks().All(ctx, s.db))
		s.Assert().Len(hooks, 1)
		s.Assert().Equal(model, hooks[0])
	})
}

func (s *hookRepoTestSuite) TestDeleteHook() {
	ctx := context.TODO()

	hookID := uuid.New()
	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		hook := models.Hook{
			HookID: hookID,
			URL:    "http://example.com",
			Method: "POST",
			Body: types.JSONB{
				"testKey": "testValue",
			},
		}
		return hook.Insert(ctx, s.db, boil.Infer())
	}, nil)
	lo.Must(0, err)

	s.Run("success", func() {
		err = db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			return s.repo.DeleteHook(ctx, tx, hookID)
		}, nil)
		s.Assert().NoError(err)
		hooks := lo.Must(models.Hooks().All(ctx, s.db))
		s.Assert().Len(hooks, 0)
	})

	s.Run("not found", func() {
		err = db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			return s.repo.DeleteHook(ctx, tx, uuid.New())
		}, nil)
		s.Assert().Error(err)
	})
}

func (s *hookRepoTestSuite) TestGetHook() {
	ctx := context.TODO()

	hookID := uuid.New()
	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		hook := models.Hook{
			HookID: hookID,
			URL:    "http://example.com",
			Method: "POST",
			Body: types.JSONB{
				"testKey": "testValue",
			},
		}
		return hook.Insert(ctx, s.db, boil.Infer())
	}, nil)
	lo.Must(0, err)

	s.Run("success", func() {
		err := db.ReadOnlyTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			hook, err := s.repo.GetHook(ctx, tx, hookID)
			s.NoError(err)
			s.Equal(hookID, hook.HookID)
			s.Equal("http://example.com", hook.URL)
			s.Equal("POST", hook.Method)
			s.Equal(types.JSONB{
				"testKey": "testValue",
			}, hook.Body)
			return nil
		}, nil)
		lo.Must(0, err)
	})

	s.Run("not found", func() {
		err := db.ReadOnlyTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			hook, err := s.repo.GetHook(ctx, tx, uuid.New())
			s.Assert().Error(err)
			s.Assert().Nil(hook)
			return nil
		}, nil)
		lo.Must(0, err)
	})
}

func (s *hookRepoTestSuite) TestGetAllHooks() {
	ctx := context.TODO()

	err := db.ExecTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
		for i := 0; i < 4; i++ {
			hook := models.Hook{
				DisplayName: fmt.Sprintf("hook%d", i+1),
				HookID:      uuid.New(),
				URL:         "http://example.com",
				Method:      "POST",
				Body: types.JSONB{
					"testKey": "testValue",
				},
			}
			err := hook.Insert(ctx, s.db, boil.Infer())
			lo.Must(0, err)
		}
		return nil
	}, nil)
	lo.Must(0, err)

	s.Run("success(limit/offset)", func() {
		err := db.ReadOnlyTx(ctx, s.db, func(ctx context.Context, tx *sql.Tx) error {
			hooks, err := s.repo.GetAllHooks(ctx, tx, 2, 1)
			s.NoError(err)
			s.Len(hooks, 2)
			s.Equal("hook2", hooks[0].DisplayName)
			s.Equal("hook3", hooks[1].DisplayName)
			return nil
		}, nil)
		lo.Must(0, err)
	})
}
