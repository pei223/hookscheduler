package web

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/pei223/hook-scheduler/pkg/web"
	"github.com/pkg/errors"
)

func TaskIDContext(c *gin.Context) {
	s := c.Param("taskID")
	if s == "" {
		web.HandleError(c, errors.New("taskId is empty"))
		return
	}
	taskID, err := uuid.Parse(s)
	if err != nil {
		web.HandleError(c, errors.New("taskId is empty"))
		return
	}
	ctx := c.Request.Context()
	ctx = task.WithTaskIDContext(ctx, taskID)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
