package web

import (
	"github.com/go-chi/chi"
	"github.com/pei223/hook-scheduler/internal/task"
	"github.com/pei223/hook-scheduler/pkg/web"
)

func NewRouter(
	taskHandler *task.TaskWebHandler,
) *chi.Mux {
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/tasks", func(r chi.Router) {
			r.Route("/{taskId}", func(r chi.Router) {
				r.Use(TaskIDContext)
				r.Get("/", web.ToHandlerFunc(taskHandler.GetTask))
			})
		})
	})
	return router
}
