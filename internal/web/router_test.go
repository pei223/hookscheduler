package web_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/pei223/hook-scheduler/internal/task/mock_task"
	"github.com/pei223/hook-scheduler/internal/web"
	"github.com/pei223/hook-scheduler/pkg/common"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
)

type routerTestSuite struct {
	suite.Suite
	router     *gin.Engine
	taskModule *mock_task.MockTaskMod
}

func (s *routerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	gomock := gomock.NewController(s.T())
	logger := common.NewLogger(context.Background(), "debug")
	s.taskModule = mock_task.NewMockTaskMod(gomock)
	s.router = web.NewRouter(
		task.NewTaskWebHandler(&logger, s.taskModule),
	)
}

func TestTaskModSuite(t *testing.T) {
	suite.Run(t, new(routerTestSuite))
}

func mustToJson(v any) string {
	b := lo.Must(json.Marshal(v))
	return string(b)
}

func mustToBody(v any) io.Reader {
	b := lo.Must(json.Marshal(v))
	return bytes.NewBuffer(b)
}
