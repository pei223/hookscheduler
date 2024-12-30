package task_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/pei223/hook-scheduler/pkg/common"
	"github.com/pei223/hook-scheduler/pkg/db"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type taskRepoTestSuite struct {
	suite.Suite
	db     *sql.DB
	repo   task.TaskRepo
	logger zerolog.Logger
}

func (s *taskRepoTestSuite) SetupSuite() {
	db := lo.Must(sql.Open("postgres", test_common.TestDatabaseConnectionString))
	s.logger = common.NewLogger(context.Background(), "debug")
	s.db = db
	s.repo = task.NewTaskRepo(db, &s.logger)
}

func (s *taskRepoTestSuite) SetupTest() {
	ctx := context.TODO()
	logger := common.NewLogger(context.Background(), "debug")
	err := db.ExecTx(ctx, &logger, s.db, func(ctx context.Context, tx *sql.Tx) error {
		_, err := models.Tasks().DeleteAll(ctx, s.db)
		return err
	})
	if err != nil {
		panic(err)
	}
}

func (s *taskRepoTestSuite) TearDownSuite() {
	s.db.Close()
}

func TestTaskRepoSuite(t *testing.T) {
	suite.Run(t, new(taskRepoTestSuite))
}

func (s *taskRepoTestSuite) TestGetTask() {
	ctx := context.TODO()

	taskID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
	err := db.ExecTx(ctx, &s.logger, s.db, func(ctx context.Context, tx *sql.Tx) error {
		task := models.Task{
			TaskID:   taskID,
			TaskName: "test",
		}
		return task.Insert(ctx, tx, boil.Infer())
	})
	if err != nil {
		panic(err)
	}

	s.Run("success", func() {
		task, err := s.repo.GetTask(ctx, taskID)
		s.NoError(err)
		s.Equal(task.TaskName, "test")
	})

	s.Run("not found", func() {
		task, err := s.repo.GetTask(ctx, uuid.MustParse("87654321-4321-8765-4321-876543218765"))
		s.Error(err)
		s.Nil(task)
	})
}
