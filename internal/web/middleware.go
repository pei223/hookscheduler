package web

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/task"
)

func TaskIDContext(c *gin.Context) {
	s := c.Param("taskID")
	if s == "" {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "taskId is empty"})
		return
	}
	taskID, err := uuid.Parse(s)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "taskId is not UUID"})
		return
	}
	ctx := c.Request.Context()
	ctx = task.WithTaskIDContext(ctx, taskID)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
