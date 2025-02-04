package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	pkgLogger "github.com/pei223/hook-scheduler/pkg/logger"
	"github.com/pei223/hook-scheduler/pkg/web"
	"github.com/rs/zerolog"
)

func generateLoggerWithRequestIDContext(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		requestID := uuid.New()
		// TODO なぜか同じ値になるので後で見直し
		logger = logger.With().Stringer("requestId", requestID).Logger()
		ctx = pkgLogger.WithContext(ctx, logger)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func NewRouter(
	hookHandler *HookRouter,
	logger zerolog.Logger,
) *gin.Engine {
	router := gin.New()

	v1 := router.Group("/api/v1")
	v1.Use(generateLoggerWithRequestIDContext(logger))
	{
		hooksRoute := v1.Group("/hooks")
		{
			hookRoute := hooksRoute.Group("/:hookID")
			{
				hookRoute.Use(hookIDContext)
				hookRoute.GET("", web.ToHandlerFunc(hookHandler.GetHook))
				hookRoute.DELETE("", web.ToHandlerFunc(hookHandler.DeleteHook))
			}
		}
		hooksRoute.POST("", web.ToHandlerFunc(hookHandler.CreateHook))
	}

	return router
}
