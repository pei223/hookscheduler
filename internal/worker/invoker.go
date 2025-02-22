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

func (i *Invoker) Start(ctx context.Context) error {
	logger := logger.FromContext(ctx)

	logger.Info().Msg("Invoker started")

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			logger.Info().Msg("Kick scheduled hooks")
			err := i.hookExecUsecase.ExecuteScheduledHooks(ctx)
			if err != nil {
				logger.Warn().Err(err).Msg("Failed to execute scheduled hooks")
			}
			// TODO: 続行不能エラーは止めたい
		case <-ctx.Done():
			logger.Info().Msg("Invoker stopped by context")
			return context.Canceled
		}
	}
}
