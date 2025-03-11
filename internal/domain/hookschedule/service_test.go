package hookschedule_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	_ "github.com/lib/pq"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule/mock_hookschedule"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
)

type hookScheduleServiceTestSuite struct {
	suite.Suite

	mockRepo *mock_hookschedule.MockHookScheduleRepoIF
	svc      *hookschedule.HookScheduleService
}

func (s *hookScheduleServiceTestSuite) SetupSuite() {
	db := lo.Must(sql.Open("postgres", test_common.TestDatabaseConnectionString))
	gomock := gomock.NewController(s.T())
	s.mockRepo = mock_hookschedule.NewMockHookScheduleRepoIF(gomock)
	s.svc = hookschedule.NewHookScheduleService(db, s.mockRepo)
}

func TestHookModSuite(t *testing.T) {
	suite.Run(t, new(hookScheduleServiceTestSuite))
}

func (s *hookScheduleServiceTestSuite) TestGetHookSchedule() {
	s.Run("success", func() {
		ctx := context.Background()
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		hookScheduleID := uuid.MustParse("22345678-1234-5678-1234-567812345678")
		mockHookSchedule := &models.HookSchedule{
			HookScheduleID: hookScheduleID,
			HookID:         hookID,
			DisplayName:    "test",
		}
		s.mockRepo.EXPECT().GetHookSchedule(gomock.Any(), gomock.Any(), hookScheduleID, true).Return(mockHookSchedule, nil).Times(1)
		hookSchedule, err := s.svc.GetHookSchedule(ctx, hookScheduleID, true)
		s.Require().NoError(err)
		s.Require().NotNil(hookSchedule)
		s.Assert().Equal(mockHookSchedule, hookSchedule)
	})
}

func (s *hookScheduleServiceTestSuite) TestDeleteHookSchedule() {
	ctx := context.TODO()

	s.Run("success", func() {
		hookScheduleID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockRepo.EXPECT().DeleteHookSchedule(gomock.Any(), gomock.Any(), hookScheduleID).Return(nil).Times(1)
		err := s.svc.DeleteHookSchedule(ctx, hookScheduleID)
		s.Require().NoError(err)
	})

	s.Run("error", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.mockRepo.EXPECT().DeleteHookSchedule(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("testerror")).Times(1)
		err := s.svc.DeleteHookSchedule(ctx, hookID)
		s.Require().Error(err)
	})
}

func (s *hookScheduleServiceTestSuite) TestCreateHookSchedule() {
	ctx := context.TODO()

	s.Run("success", func() {
		hookID := uuid.New()
		params := &hookschedule.HookScheduleCreateParams{
			DisplayName: "test",
		}
		s.mockRepo.EXPECT().CreateHookSchedule(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.HookSchedule{
			HookID:      hookID,
			DisplayName: params.DisplayName,
		}, nil).Times(1)
		hook, err := s.svc.CreateHookSchedule(ctx, params)
		s.Require().NoError(err)
		s.Require().NotNil(hook)
		s.Assert().Equal(*hook, models.HookSchedule{
			HookID:      hookID,
			DisplayName: params.DisplayName,
		})
	})

	s.Run("error", func() {
		params := &hookschedule.HookScheduleCreateParams{}
		s.mockRepo.EXPECT().CreateHookSchedule(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("testerror")).Times(1)
		hook, err := s.svc.CreateHookSchedule(ctx, params)
		s.Require().Error(err)
		s.Require().Nil(hook)
	})
}

func (s *hookScheduleServiceTestSuite) TestListHookSchedules() {
	ctx := context.TODO()

	s.Run("success", func() {
		s.mockRepo.EXPECT().ListHookSchedules(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(models.HookScheduleSlice{
			&models.HookSchedule{
				HookScheduleID: uuid.New(),
				DisplayName:    "test1",
			},
			&models.HookSchedule{
				HookScheduleID: uuid.New(),
				DisplayName:    "test2",
			},
		}, 3, nil).Times(1)
		hooks, total, err := s.svc.ListHookSchedules(ctx, &types.ListParams{}, false)
		s.Require().NoError(err)
		s.Require().NotNil(hooks)
		s.Assert().Len(hooks, 2)
		s.Assert().Equal(3, total)
	})

	s.Run("error", func() {
		s.mockRepo.EXPECT().ListHookSchedules(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, 0, errors.New("test err")).Times(1)
		hooks, _, err := s.svc.ListHookSchedules(ctx, &types.ListParams{}, false)
		s.Require().Error(err)
		s.Require().Nil(hooks)
	})
}
