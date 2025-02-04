package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pei223/hook-scheduler/internal/domain/hook"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/usecase"
	"github.com/pei223/hook-scheduler/internal/usecase/mock_usecase"
	"github.com/pei223/hook-scheduler/pkg/common"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
)

type hookModTestSuite struct {
	suite.Suite

	mockHookService *mock_usecase.MockHookServiceIF
	hookUsecase     *usecase.HookUsecase
	logger          zerolog.Logger
}

func (s *hookModTestSuite) SetupSuite() {
	s.logger = common.NewLogger(context.Background(), "debug")
	gomock := gomock.NewController(s.T())
	s.mockHookService = mock_usecase.NewMockHookServiceIF(gomock)
	s.hookUsecase = usecase.NewHookUsecase(s.mockHookService)
}

func TestHookModSuite(t *testing.T) {
	suite.Run(t, new(hookModTestSuite))
}

func (s *hookModTestSuite) TestGetHook() {
	s.Run("success", func() {
		ctx := context.Background()
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		mockHook := &models.Hook{
			HookID:      hookID,
			DisplayName: "test",
			URL:         "http://test.com",
			Method:      "POST",
		}
		s.mockHookService.EXPECT().GetHook(gomock.Any(), gomock.Any()).Return(mockHook, nil).Times(1)
		hook, err := s.hookUsecase.GetHook(ctx, hookID)
		s.Require().NoError(err)
		s.Require().NotNil(hook)
		s.Assert().Equal(mockHook, hook)
	})

	s.Run("error", func() {
		ctx := context.Background()
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockHookService.EXPECT().GetHook(gomock.Any(), gomock.Any()).Return(nil, errors.New("testerror")).Times(1)
		hook, err := s.hookUsecase.GetHook(ctx, hookID)
		s.Require().Error(err)
		s.Require().Nil(hook)
	})
}

func (s *hookModTestSuite) TestDeleteHook() {
	ctx := context.TODO()

	s.Run("success", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockHookService.EXPECT().DeleteHook(gomock.Any(), gomock.Any()).Return(nil).Times(1)
		err := s.hookUsecase.DeleteHook(ctx, hookID)
		s.Require().NoError(err)
	})

	s.Run("error", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockHookService.EXPECT().DeleteHook(gomock.Any(), gomock.Any()).Return(errors.New("testerror")).Times(1)
		err := s.hookUsecase.DeleteHook(ctx, hookID)
		s.Require().Error(err)
	})
}

func (s *hookModTestSuite) TestCreateHook() {
	ctx := context.TODO()

	s.Run("success", func() {
		hookID := uuid.New()
		params := &hook.HookCreateParams{
			DisplayName: "test",
			URL:         "http://test.com",
			Method:      "POST",
		}
		s.mockHookService.EXPECT().CreateHook(gomock.Any(), gomock.Any()).Return(&models.Hook{
			HookID:      hookID,
			DisplayName: params.DisplayName,
			URL:         params.URL,
			Method:      params.Method,
		}, nil).Times(1)
		hook, err := s.hookUsecase.CreateHook(ctx, params)
		s.Assert().NoError(err)
		s.Assert().NotNil(hook)
		s.Assert().Equal(*hook, models.Hook{
			HookID:      hookID,
			DisplayName: params.DisplayName,
			URL:         params.URL,
			Method:      params.Method,
		})
	})

	s.Run("error", func() {
		params := &hook.HookCreateParams{
			DisplayName: "test",
			URL:         "http://test.com",
			Method:      "POST",
		}
		s.mockHookService.EXPECT().CreateHook(gomock.Any(), gomock.Any()).Return(nil, errors.New("testerror")).Times(1)
		hook, err := s.hookUsecase.CreateHook(ctx, params)
		s.Assert().Error(err)
		s.Assert().Nil(hook)
	})
}
