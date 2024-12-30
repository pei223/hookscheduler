package task_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/pei223/hook-scheduler/internal/task/mock_task"
	"github.com/stretchr/testify/suite"
)

type taskModTestSuite struct {
	suite.Suite

	repo *mock_task.MockTaskRepo
	mod  task.TaskMod
}

func (s *taskModTestSuite) SetupSuite() {
	gomock := gomock.NewController(s.T())
	s.repo = mock_task.NewMockTaskRepo(gomock)
	s.mod = task.NewTaskMod(s.repo)
}

func TestTaskModSuite(t *testing.T) {
	suite.Run(t, new(taskModTestSuite))
}

func (s *taskModTestSuite) TestGetTask() {
	s.Run("success", func() {
		ctx := context.Background()
		taskID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.repo.EXPECT().GetTask(gomock.Any(), gomock.Any()).Return(&models.Task{
			TaskID:   taskID,
			TaskName: "test",
		}, nil).Times(1)
		task, err := s.mod.GetTask(ctx, taskID)
		s.Require().NoError(err)
		s.Require().NotNil(task)
	})
}
