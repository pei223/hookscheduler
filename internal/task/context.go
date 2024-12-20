package task

import (
	"context"

	"github.com/google/uuid"
)

type taskIDContextKey struct{}

func WithTaskIDContext(ctx context.Context, ItemId uuid.UUID) context.Context {
	return context.WithValue(ctx, taskIDContextKey{}, ItemId)
}

func TaskIDFromContext(ctx context.Context) uuid.UUID {
	return ctx.Value(taskIDContextKey{}).(uuid.UUID)
}
