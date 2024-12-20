package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/task"
)

func TaskIDContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := chi.URLParam(r, "taskId")
		if s == "" {
			// renderer.RenderError(w, apperrors.NewBadRequestErr("ItemId is empty", []apperrors.InvalidParam{}))
			return
		}
		taskId, err := uuid.Parse(s)
		if err != nil {
			// renderer.RenderError(w, apperrors.NewBadRequestErr("ItemId is not UUID", []apperrors.InvalidParam{}))
			return
		}
		ctx := task.WithTaskIDContext(r.Context(), taskId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
