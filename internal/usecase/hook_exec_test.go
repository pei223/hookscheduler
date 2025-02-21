package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/pei223/hook-scheduler/internal/usecase"
	"github.com/pei223/hook-scheduler/internal/usecase/mock_usecase"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
)

type hookExecTestSuite struct {
	suite.Suite

	mockHookExecService *mock_usecase.MockHookExecServiceIF
	hookExecUsecase     *usecase.HookExecUsecase
}

func (s *hookExecTestSuite) SetupSuite() {
	db := lo.Must(sql.Open("postgres", test_common.TestDatabaseConnectionString))
	gomock := gomock.NewController(s.T())
	s.mockHookExecService = mock_usecase.NewMockHookExecServiceIF(gomock)
	s.hookExecUsecase = usecase.NewHookExecUsecase(db, s.mockHookExecService)
}

func TestHookExecSuite(t *testing.T) {
	suite.Run(t, new(hookExecTestSuite))
}

func (s *hookExecTestSuite) TestExecScheduledHook() {
	ctx := context.Background()

	s.Run("success", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockHookExecService.EXPECT().GetAllHooks(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(models.HookSlice{
			{
				HookID:      hookID,
				DisplayName: "test",
				URL:         "http://test1.com",
				Method:      "POST",
			},
			{
				HookID:      hookID,
				DisplayName: "test2",
				URL:         "http://test2.com",
				Method:      "POST",
			},
		}, nil).Times(1)
		s.mockHookExecService.EXPECT().GetAllHooks(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(models.HookSlice{}, nil).Times(1)
		s.mockHookExecService.EXPECT().ExecHookInTx(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(200, nil).Times(2)
		err := s.hookExecUsecase.ExecuteScheduledHooks(ctx)
		s.Require().NoError(err)
	})

	s.Run("success (execute hook error)", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockHookExecService.EXPECT().GetAllHooks(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(models.HookSlice{
			{
				HookID:      hookID,
				DisplayName: "test",
				URL:         "http://test1.com",
				Method:      "POST",
			},
			{
				HookID:      hookID,
				DisplayName: "test2",
				URL:         "http://test2.com",
				Method:      "POST",
			},
		}, nil).Times(1)
		s.mockHookExecService.EXPECT().GetAllHooks(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(models.HookSlice{}, nil).Times(1)
		s.mockHookExecService.EXPECT().ExecHookInTx(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(0, errors.New("execute hook error")).Times(1)
		s.mockHookExecService.EXPECT().ExecHookInTx(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(200, nil).Times(1)
		err := s.hookExecUsecase.ExecuteScheduledHooks(ctx)
		s.Require().NoError(err)
	})

	s.Run("error at get hooks", func() {
		s.mockHookExecService.EXPECT().GetAllHooks(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(nil, errors.New("testErr")).Times(1)
		err := s.hookExecUsecase.ExecuteScheduledHooks(ctx)
		s.Require().Error(err)
	})

	s.Run("error at context canceled", func() {
		ctx, canceler := context.WithCancel(context.Background())
		canceler()
		err := s.hookExecUsecase.ExecuteScheduledHooks(ctx)
		s.Require().Error(err)
		s.Require().ErrorIs(err, context.Canceled)
	})
}

func (s *hookExecTestSuite) TestExecHook() {
	s.Run("success", func() {
		ctx := context.Background()
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockHookExecService.EXPECT().ExecHookInTx(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(200, nil).Times(1)
		err := s.hookExecUsecase.ExecHook(ctx, hookID)
		s.Require().NoError(err)
	})

	s.Run("error", func() {
		ctx := context.Background()
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockHookExecService.EXPECT().ExecHookInTx(
			gomock.Any(),
			gomock.Any(),
			gomock.Any(),
		).Return(0, errors.New("exec hook error")).Times(1)
		err := s.hookExecUsecase.ExecHook(ctx, hookID)
		s.Require().Error(err)
	})
}
