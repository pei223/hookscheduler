package webapi_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/types"
	"github.com/samber/lo"
)

func (s *routerTestSuite) TestGetHook() {
	s.Run("success", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookUsecase.EXPECT().GetHook(gomock.Any(), hookID).Return(&models.Hook{
			HookID:      hookID,
			DisplayName: "test",
			URL:         "http://test.com",
			Method:      "POST",
			Body: types.JSONB{
				"gettestkey": "getvalue",
				"testlist": []string{
					"test1", "test2",
				},
			},
			Headers: types.JSONB{
				"gettestkey": "getvalue",
				"testlist": []string{
					"test1", "test2",
				},
			},
		}, nil).Times(1)

		req := lo.Must(http.NewRequest("GET", "/api/v1/hooks/"+hookID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusOK, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})

	s.Run("not found", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookUsecase.EXPECT().GetHook(gomock.Any(), hookID).Return(nil, sql.ErrNoRows).Times(1)

		req := lo.Must(http.NewRequest("GET", "/api/v1/hooks/"+hookID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusNotFound, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
}

func (s *routerTestSuite) TestDeleteHook() {
	s.Run("success", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookUsecase.EXPECT().DeleteHook(gomock.Any(), hookID).Return(nil).Times(1)

		req := lo.Must(http.NewRequest("DELETE", "/api/v1/hooks/"+hookID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusNoContent, w.Code)
	})

	s.Run("not found", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookUsecase.EXPECT().DeleteHook(gomock.Any(), hookID).Return(sql.ErrNoRows).Times(1)

		req := lo.Must(http.NewRequest("DELETE", "/api/v1/hooks/"+hookID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusNotFound, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
}

func (s *routerTestSuite) TestCreateHooks() {
	dummyBody := types.JSONB{
		"createtestkey": "createvalue",
		"testlist": []string{
			"test1", "test2",
		},
	}
	dummyHeaders := types.JSONB{
		"createtestkey": "createvalue",
		"testlist": []string{
			"test1", "test2",
		},
	}
	s.Run("success", func() {
		hookID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.hookUsecase.EXPECT().CreateHook(gomock.Any(), gomock.Any()).Return(&models.Hook{
			HookID:      hookID,
			DisplayName: "test",
			URL:         "http://test.com",
			Method:      "POST",
			Body:        dummyBody,
			Headers:     dummyHeaders,
		}, nil).Times(1)

		req := lo.Must(http.NewRequest("POST", "/api/v1/hooks", mustToBody(
			map[string]any{
				"displayName": "test",
				"method":      "GET",
				"url":         "http://test.test",
				"body":        dummyBody,
				"headers":     dummyHeaders,
			},
		)))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusCreated, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
	s.Run("contains invalid params", func() {
		req := lo.Must(http.NewRequest("POST", "/api/v1/hooks", mustToBody(
			map[string]any{
				"displayName": "1234567890123456789012345678901234567890",
				"url":         "http://test.test",
				"method":      "GET",
				"body":        "test",
				"headers":     dummyHeaders,
			},
		)))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusBadRequest, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
	s.Run("contains empty field", func() {
		req := lo.Must(http.NewRequest("POST", "/api/v1/hooks", mustToBody(
			map[string]any{
				"url":     "http://test.test",
				"method":  "GET",
				"body":    dummyBody,
				"headers": dummyHeaders,
			},
		)))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		s.Assert().Equal(http.StatusBadRequest, w.Code)
		snaps.MatchJSON(s.T(), w.Body.String())
	})
}
