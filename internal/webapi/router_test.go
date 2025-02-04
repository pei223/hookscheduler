package webapi_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/pei223/hook-scheduler/internal/webapi"
	"github.com/pei223/hook-scheduler/internal/webapi/mock_webapi"
	"github.com/pei223/hook-scheduler/pkg/logger"
	"github.com/samber/lo"
	"github.com/stretchr/testify/suite"
)

type routerTestSuite struct {
	suite.Suite
	router      *gin.Engine
	hookUsecase *mock_webapi.MockHookUsecaseIF
}

func (s *routerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	gomock := gomock.NewController(s.T())
	s.hookUsecase = mock_webapi.NewMockHookUsecaseIF(gomock)
	s.router = webapi.NewRouter(
		webapi.NewHookRouter(s.hookUsecase),
		logger.NewLogger(context.TODO(), "debug"),
	)
}

func TestHookModSuite(t *testing.T) {
	suite.Run(t, new(routerTestSuite))
}

func mustToBody(v any) io.Reader {
	b := lo.Must(json.Marshal(v))
	return bytes.NewBuffer(b)
}
