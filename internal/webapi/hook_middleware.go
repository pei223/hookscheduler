package webapi

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/pkg/web"
	"github.com/pkg/errors"
)

type hookIDContextKey struct{}

func hookIDContext(c *gin.Context) {
	s := c.Param("hookID")
	if s == "" {
		web.HandleError(c, errors.New("hookId is empty"))
		return
	}
	hookID, err := uuid.Parse(s)
	if err != nil {
		web.HandleError(c, errors.New("hookId is empty"))
		return
	}
	ctx := c.Request.Context()
	ctx = WithHookIDContext(ctx, hookID)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func WithHookIDContext(ctx context.Context, hookID uuid.UUID) context.Context {
	return context.WithValue(ctx, hookIDContextKey{}, hookID)
}

func HookIDFromContext(ctx context.Context) uuid.UUID {
	return ctx.Value(hookIDContextKey{}).(uuid.UUID)
}
