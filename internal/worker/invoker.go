package worker

import (
	"context"
	"time"

	"github.com/pei223/hook-scheduler/pkg/logger"
)

type HookExecUsecase interface {
	ExecuteScheduledHooks(ctx context.Context) error
}

type Invoker struct {
	hookExecUsecase HookExecUsecase
}

func NewInvoker(hookExecUsecase HookExecUsecase) *Invoker {
	return &Invoker{
		hookExecUsecase: hookExecUsecase,
	}
}

func (i *Invoker) Start(ctx context.Context) {
	logger := logger.FromContext(ctx)

	logger.Info().Msg("Invoker started")

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			logger.Info().Msg("Kick scheduled hooks")
			i.hookExecUsecase.ExecuteScheduledHooks(ctx)
		case <-ctx.Done():
			logger.Info().Msg("Invoker stopped by context")
			return
		}
	}
}
