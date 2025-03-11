package webapi

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/pkg/web"
	"github.com/pkg/errors"
)

type hookScheduleIDContextKey struct{}

func hookScheduleIDContext(c *gin.Context) {
	s := c.Param("hookScheduleID")
	if s == "" {
		web.HandleError(c, errors.New("hookScheduleId is empty"))
		return
	}
	hookScheduleID, err := uuid.Parse(s)
	if err != nil {
		web.HandleError(c, errors.New("hookScheduleId is empty"))
		return
	}
	ctx := c.Request.Context()
	ctx = WithHookScheduleIDContext(ctx, hookScheduleID)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func WithHookScheduleIDContext(ctx context.Context, hookScheduleID uuid.UUID) context.Context {
	return context.WithValue(ctx, hookScheduleIDContextKey{}, hookScheduleID)
}

func HookScheduleIDFromContext(ctx context.Context) uuid.UUID {
	return ctx.Value(hookScheduleIDContextKey{}).(uuid.UUID)
}
