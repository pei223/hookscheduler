package usecase_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/test_common"
	"github.com/pei223/hook-scheduler/internal/usecase"
	mock_usecase "github.com/pei223/hook-scheduler/internal/usecase/mock_usecase"
	"github.com/pei223/hook-scheduler/pkg/types"
)

type hookScheduleUseCaseTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	mockSvc *mock_usecase.MockHookScheduleServiceIF
	uc      *usecase.HookScheduleUsecase
}

func (s *hookScheduleUseCaseTestSuite) SetupTest() {
	db := lo.Must(sql.Open("postgres", test_common.TestDatabaseConnectionString))
	s.ctrl = gomock.NewController(s.T())
	s.mockSvc = mock_usecase.NewMockHookScheduleServiceIF(s.ctrl)
	s.uc = usecase.NewHookScheduleUsecase(db, s.mockSvc)
}

func TestHookScheduleUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(hookScheduleUseCaseTestSuite))
}

func (s *hookScheduleUseCaseTestSuite) TestCreateHookSchedule() {
	s.Run("success", func() {
		hookScheduleID := uuid.New()
		hookID := uuid.New()
		hookSchedule := models.HookSchedule{
			HookScheduleID: hookScheduleID,
			HookID:         hookID,
			DisplayName:    "test",
		}
		createParams := &hookschedule.HookScheduleCreateParams{
			HookID:      hookID,
			DisplayName: "test",
		}
		s.mockSvc.EXPECT().CreateHookSchedule(gomock.Any(), createParams).Return(&hookSchedule, nil)

		result, err := s.uc.CreateHookSchedule(context.Background(), createParams)
		require.NoError(s.T(), err)
		assert.Equal(s.T(), &hookSchedule, result)
	})
	s.Run("error", func() {
		createParams := &hookschedule.HookScheduleCreateParams{}
		s.mockSvc.EXPECT().CreateHookSchedule(gomock.Any(), createParams).Return(nil, errors.New("create error"))

		result, err := s.uc.CreateHookSchedule(context.Background(), createParams)
		assert.Nil(s.T(), result)
		assert.Error(s.T(), err)
	})
}

func (s *hookScheduleUseCaseTestSuite) TestDeleteHookSchedule() {
	s.Run("success", func() {
		hookScheduleID := uuid.New()
		s.mockSvc.EXPECT().DeleteHookSchedule(gomock.Any(), hookScheduleID).Return(nil)

		err := s.uc.DeleteHookSchedule(context.Background(), hookScheduleID)
		require.NoError(s.T(), err)
	})
	s.Run("error", func() {
		hookScheduleID := uuid.New()
		s.mockSvc.EXPECT().DeleteHookSchedule(gomock.Any(), hookScheduleID).Return(errors.New("delete error"))

		err := s.uc.DeleteHookSchedule(context.Background(), hookScheduleID)
		require.Error(s.T(), err)
	})
}

func (s *hookScheduleUseCaseTestSuite) TestGetHookSchedule() {
	s.Run("success", func() {
		hookScheduleID := uuid.New()
		hookID := uuid.New()
		hookSchedule := models.HookSchedule{
			HookScheduleID: hookScheduleID,
			HookID:         hookID,
			DisplayName:    "test",
		}
		s.mockSvc.EXPECT().GetHookSchedule(gomock.Any(), hookScheduleID, false).Return(&hookSchedule, nil)

		result, err := s.uc.GetHookSchedule(context.Background(), hookScheduleID, false)
		require.NoError(s.T(), err)
		assert.Equal(s.T(), &hookSchedule, result)
	})
	s.Run("error", func() {
		hookScheduleID := uuid.New()
		s.mockSvc.EXPECT().GetHookSchedule(gomock.Any(), hookScheduleID, true).Return(nil, errors.New("get error"))

		result, err := s.uc.GetHookSchedule(context.Background(), hookScheduleID, true)
		assert.Nil(s.T(), result)
		assert.Error(s.T(), err)
	})
}

func (s *hookScheduleUseCaseTestSuite) TestListHookSchedules() {
	s.Run("success", func() {
		hookScheduleID := uuid.New()
		hookID := uuid.New()
		hookSchedule := models.HookSchedule{
			HookScheduleID: hookScheduleID,
			HookID:         hookID,
			DisplayName:    "test",
		}
		s.mockSvc.EXPECT().ListHookSchedules(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.HookScheduleSlice{&hookSchedule}, 3, nil)

		result, total, err := s.uc.ListHookSchedules(context.Background(), &types.ListParams{}, false)
		require.NoError(s.T(), err)
		assert.Equal(s.T(), models.HookScheduleSlice{&hookSchedule}, result)
		assert.Equal(s.T(), 3, total)
	})
	s.Run("error", func() {
		s.mockSvc.EXPECT().ListHookSchedules(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, 0, errors.New("list error"))
		result, _, err := s.uc.ListHookSchedules(context.Background(), &types.ListParams{}, false)
		assert.Nil(s.T(), result)
		assert.Error(s.T(), err)
	})
}
