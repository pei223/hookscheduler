package usecase

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/db"
	"github.com/pei223/hook-scheduler/pkg/logger"
)

type HookExecServiceIF interface {
	GetAllHooks(ctx context.Context, limit, offset int) (models.HookSlice, error)
	ExecHookInTx(ctx context.Context, tx *sql.Tx, hookID uuid.UUID) (int, error)
}

type HookExecUsecase struct {
	hookExecSvc HookExecServiceIF
	db          *sql.DB
}

func NewHookExecUsecase(db *sql.DB, hookExecSvc HookExecServiceIF) *HookExecUsecase {
	return &HookExecUsecase{
		hookExecSvc: hookExecSvc,
		db:          db,
	}
}

func (t *HookExecUsecase) ExecuteScheduledHooks(ctx context.Context) error {
	limit := 100
	offset := 0
	logger := logger.FromContext(ctx)
	for {
		select {
		case <-ctx.Done():
			return context.Canceled
		default:
		}
		// TODO ここで正しくスケジュールに適した処理を取得する
		hooks, err := t.hookExecSvc.GetAllHooks(ctx, limit, offset)
		if err != nil {
			return fmt.Errorf("failed to get hooks: %w", err)
		}
		offset += limit
		if len(hooks) == 0 {
			logger.Info().Msg("no hooks to execute")
			return nil
		}
		// TODO ここでgoroutine並列処理
		for _, hook := range hooks {
			err := t.ExecHook(ctx, hook.HookID)
			// ignore err to continue other processing
			if err != nil {
				logger.Error().Err(err).Msg("failed to execute hook")
			}
		}
	}
}

func (t *HookExecUsecase) ExecHook(ctx context.Context, hookID uuid.UUID) error {
	_logger := logger.FromContext(ctx).With().Stringer("hookId", hookID).Logger()
	err := db.ExecTx(ctx, t.db, func(ctx context.Context, tx *sql.Tx) error {
		status, err := t.hookExecSvc.ExecHookInTx(logger.WithContext(ctx, _logger), tx, hookID)
		if err != nil {
			return fmt.Errorf("failed to execute hook in usecase: %w", err)
		}
		_logger.Info().Int("status", status).Msg("hook executed")
		// TODO 履歴書き込み
		return nil
	}, nil)
	if err != nil {
		_logger.Error().Err(err).Msg("failed to execute hook")
		return fmt.Errorf("failed to execute hook in tx: %w", err)
	}
	return nil
}
