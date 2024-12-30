package web

import (
	"github.com/gin-gonic/gin"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/pei223/hook-scheduler/pkg/web"
)

func NewRouter(
	taskHandler *task.TaskWebHandler,
) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		tasksRoute := v1.Group("/tasks")
		{
			taskRoute := tasksRoute.Group("/:taskID")
			{
				taskRoute.Use(TaskIDContext)
				taskRoute.GET("", web.ToHandlerFunc(taskHandler.GetTask))
			}
		}
	}

	return router
}
