package webapi_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hookschedule"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/samber/lo"
)

func (s *routerTestSuite) TestGetHookSchedule() {
	s.Run("success", func() {
		hookScheduleID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookScheduleUsecase.EXPECT().GetHookSchedule(gomock.Any(), hookScheduleID, false).Return(&models.HookSchedule{
			HookScheduleID:        hookScheduleID,
			HookID:                hookID,
			DisplayName:           "get test",
			Description:           "get test",
			ScheduleFrequencyUnit: "every_minute",
			ScheduleTimeMonth:     2,
			ScheduleTimeDay:       3,
			ScheduleTimeHour:      4,
			ScheduleTimeMinute:    5,
			ScheduleTimeSecond:    6,
		}, nil).Times(1)

		req := lo.Must(http.NewRequest("GET", "/api/v1/hook-schedules/"+hookScheduleID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusOK, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})

	s.Run("not found", func() {
		hookScheduleID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookScheduleUsecase.EXPECT().GetHookSchedule(gomock.Any(), hookScheduleID, false).Return(nil, sql.ErrNoRows).Times(1)

		req := lo.Must(http.NewRequest("GET", "/api/v1/hook-schedules/"+hookScheduleID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusNotFound, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
}

func (s *routerTestSuite) TestDeleteHookSchedule() {
	s.Run("success", func() {
		hookScheduleID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookScheduleUsecase.EXPECT().DeleteHookSchedule(gomock.Any(), hookScheduleID).Return(nil).Times(1)

		req := lo.Must(http.NewRequest("DELETE", "/api/v1/hook-schedules/"+hookScheduleID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusNoContent, w.Code)
	})

	s.Run("not found", func() {
		hookScheduleID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookScheduleUsecase.EXPECT().DeleteHookSchedule(gomock.Any(), hookScheduleID).Return(sql.ErrNoRows).Times(1)

		req := lo.Must(http.NewRequest("DELETE", "/api/v1/hook-schedules/"+hookScheduleID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusNotFound, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
}

func (s *routerTestSuite) TestCreateHookSchedule() {
	s.Run("success", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		hookScheduleID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		params := &hookschedule.HookScheduleCreateParams{
			HookID:                hookID,
			DisplayName:           "create test",
			Description:           "create test",
			ScheduleFrequencyUnit: "every_minute",
			ScheduleTimeMonth:     5,
			ScheduleTimeDay:       4,
			ScheduleTimeHour:      3,
			ScheduleTimeMinute:    2,
			ScheduleTimeSecond:    1,
		}
		s.hookScheduleUsecase.EXPECT().CreateHookSchedule(gomock.Any(), params).Return(&models.HookSchedule{
			HookScheduleID:        hookScheduleID,
			HookID:                hookID,
			DisplayName:           "create test",
			Description:           "create test",
			ScheduleFrequencyUnit: "every_minute",
			ScheduleTimeMonth:     5,
			ScheduleTimeDay:       4,
			ScheduleTimeHour:      3,
			ScheduleTimeMinute:    2,
			ScheduleTimeSecond:    1,
		}, nil).Times(1)

		req := lo.Must(http.NewRequest("POST", "/api/v1/hook-schedules", mustToBody(
			map[string]any{
				"hookId":                hookID,
				"displayName":           "create test",
				"description":           "create test",
				"scheduleFrequencyUnit": "every_minute",
				"scheduleTimeMonth":     5,
				"scheduleTimeDay":       4,
				"scheduleTimeHour":      3,
				"scheduleTimeMinute":    2,
				"scheduleTimeSecond":    1,
			},
		)))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusCreated, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
	s.Run("contains invalid params", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		req := lo.Must(http.NewRequest("POST", "/api/v1/hook-schedules", mustToBody(
			map[string]any{
				"hookId":                hookID,
				"displayName":           "create test",
				"description":           "create test",
				"scheduleFrequencyUnit": 10000,
				"scheduleTimeMonth":     5,
				"scheduleTimeDay":       4,
				"scheduleTimeHour":      3,
				"scheduleTimeMinute":    2,
				"scheduleTimeSecond":    1,
			},
		)))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusBadRequest, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
	s.Run("contains empty field", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		req := lo.Must(http.NewRequest("POST", "/api/v1/hook-schedules", mustToBody(
			map[string]any{
				"hookId":      hookID,
				"displayName": "create test",
				"description": "create test",
			},
		)))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusBadRequest, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
}

func (s *routerTestSuite) TestListHookSchedules() {
	s.Run("success (no limit and offset)", func() {
		s.hookScheduleUsecase.EXPECT().ListHookSchedules(gomock.Any(), &types.ListParams{
			Limit:  10,
			Offset: 0,
		}, false).Return(models.HookScheduleSlice{
			{
				HookID:      uuid.MustParse("12345678-1234-5678-1234-567812345678"),
				DisplayName: "list test",
			},
		}, 4, nil).Times(1)

		req := lo.Must(http.NewRequest("GET", "/api/v1/hook-schedules", nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusOK, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
	s.Run("success (limit and offset)", func() {
		s.hookScheduleUsecase.EXPECT().ListHookSchedules(gomock.Any(), &types.ListParams{
			Limit:  10,
			Offset: 20,
		}, false).Return(models.HookScheduleSlice{
			{
				HookID:      uuid.MustParse("12345678-1234-5678-1234-567812345678"),
				DisplayName: "list test",
			},
			{
				HookID:      uuid.MustParse("12345678-1234-5678-1234-567812345678"),
				DisplayName: "list test2",
			},
		}, 3, nil).Times(1)

		req := lo.Must(http.NewRequest("GET", "/api/v1/hook-schedules?limit=10&offset=20", nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusOK, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
}
