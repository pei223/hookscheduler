package hook_test

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"
	_ "github.com/lib/pq"
	"github.com/pei223/hook-scheduler/internal/domain/hook"
	"github.com/pei223/hook-scheduler/internal/domain/hook/mock_hook"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
	"resty.dev/v3"
)

type hookModTestSuite struct {
	suite.Suite

	mockRepo *mock_hook.MockHookRepo
	svc      *hook.HookService
	client   *http.Client
}

func (s *hookModTestSuite) SetupSuite() {
	db := lo.Must(sql.Open("postgres", test_common.TestDatabaseConnectionString))
	gomock := gomock.NewController(s.T())
	s.mockRepo = mock_hook.NewMockHookRepo(gomock)
	s.client = &http.Client{}
	s.svc = hook.NewHookService(db, s.mockRepo, resty.NewWithClient(s.client))
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
		s.mockRepo.EXPECT().GetHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockHook, nil).Times(1)
		hook, err := s.svc.GetHook(ctx, hookID)
		s.Require().NoError(err)
		s.Require().NotNil(hook)
		s.Assert().Equal(mockHook, hook)
	})
}

func (s *hookModTestSuite) TestDeleteHook() {
	ctx := context.TODO()

	s.Run("success", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockRepo.EXPECT().DeleteHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)
		err := s.svc.DeleteHook(ctx, hookID)
		s.Require().NoError(err)
	})

	s.Run("error", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockRepo.EXPECT().DeleteHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("testerror")).Times(1)
		err := s.svc.DeleteHook(ctx, hookID)
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
		s.mockRepo.EXPECT().CreateHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.Hook{
			HookID:      hookID,
			DisplayName: params.DisplayName,
			URL:         params.URL,
			Method:      params.Method,
		}, nil).Times(1)
		hook, err := s.svc.CreateHook(ctx, params)
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
		s.mockRepo.EXPECT().CreateHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("testerror")).Times(1)
		hook, err := s.svc.CreateHook(ctx, params)
		s.Assert().Error(err)
		s.Assert().Nil(hook)
	})
}

func (s *hookModTestSuite) TestGetAllHooks() {
	ctx := context.TODO()

	s.Run("success", func() {
		s.mockRepo.EXPECT().GetAllHooks(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(models.HookSlice{
			&models.Hook{
				HookID:      uuid.New(),
				DisplayName: "test",
			},
			&models.Hook{
				HookID:      uuid.New(),
				DisplayName: "test2",
			},
		}, nil).Times(1)
		hooks, err := s.svc.GetAllHooks(ctx, 10, 0)
		s.Assert().NoError(err)
		s.Assert().NotNil(hooks)
		s.Assert().Len(hooks, 2)
	})

	s.Run("error", func() {
		s.mockRepo.EXPECT().GetAllHooks(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("test err")).Times(1)
		hooks, err := s.svc.GetAllHooks(ctx, 10, 0)
		s.Assert().Error(err)
		s.Assert().Nil(hooks)
	})
}

func (s *hookModTestSuite) TestExecHookInTx() {
	ctx := context.TODO()

	s.Run("success", func() {
		s.mockRepo.EXPECT().GetHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.Hook{
			DisplayName: "test",
			URL:         "http://test.com",
			Method:      "POST",
		}, nil)

		httpmock.ActivateNonDefault(s.client)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", "http://test.com", httpmock.NewStringResponder(201, ""))

		status, err := s.svc.ExecHookInTx(ctx, nil, uuid.New())
		s.Require().NoError(err)
		s.Require().Equal(201, status)
	})

	s.Run("success (internal error)", func() {
		s.mockRepo.EXPECT().GetHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.Hook{
			DisplayName: "test",
			URL:         "http://test.com",
			Method:      "POST",
		}, nil)

		httpmock.ActivateNonDefault(s.client)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("POST", "http://test.com", httpmock.NewStringResponder(500, ""))

		status, err := s.svc.ExecHookInTx(ctx, nil, uuid.New())
		s.Require().NoError(err)
		s.Require().Equal(500, status)
	})

	s.Run("get hook error", func() {
		s.mockRepo.EXPECT().GetHook(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("test error"))
		_, err := s.svc.ExecHookInTx(ctx, nil, uuid.New())
		s.Require().Error(err)
	})
}
