package web_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func (s *routerTestSuite) TestGetTasks() {
	s.Run("success", func() {
		taskID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.taskModule.EXPECT().GetTask(gomock.Any(), taskID).Return(&models.Task{
			TaskID:   taskID,
			TaskName: "test",
		}, nil).Times(1)

		req := lo.Must(http.NewRequest("GET", "/api/v1/tasks/"+taskID.String(), nil))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		assert.Equal(s.T(), http.StatusOK, w.Code)
		expected := mustToJson(task.Task{
			TaskId: taskID.String(),
			Name:   "test",
			Status: "pending",
		})
		assert.JSONEq(s.T(), expected, w.Body.String())
	})
}

func (s *routerTestSuite) TestCreateTasks() {
	s.Run("success", func() {
		taskID := uuid.MustParse("12345678-1234-5678-1234-567812345678")
		s.taskModule.EXPECT().CreateTask(gomock.Any(), gomock.Any()).Return(&models.Task{
			TaskID:   taskID,
			TaskName: "createtest",
		}, nil).Times(1)

		req := lo.Must(http.NewRequest("POST", "/api/v1/tasks", mustToBody(
			task.TaskCreateParams{
				Name: "createtest",
			},
		)))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		assert.Equal(s.T(), http.StatusCreated, w.Code)
		expected := mustToJson(task.Task{
			TaskId: taskID.String(),
			Name:   "createtest",
			Status: "pending",
		})
		assert.JSONEq(s.T(), expected, w.Body.String())
	})
	s.Run("invalid param", func() {
		req := lo.Must(http.NewRequest("POST", "/api/v1/tasks", mustToBody(
			task.TaskCreateParams{
				Name: "1234567890123456789012345678901234567890123456789012345678901234567890",
			},
		)))
		w := httptest.NewRecorder()
		s.router.ServeHTTP(w, req)

		assert.Equal(s.T(), http.StatusBadRequest, w.Code)
		assert.Contains(s.T(), w.Body.String(), "Name must be a maximum of 20 characters")
	})
}
