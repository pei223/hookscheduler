package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/domain/hook"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/logger"
)

type HookServiceIF interface {
	GetHook(ctx context.Context, hookId uuid.UUID) (*models.Hook, error)
	DeleteHook(ctx context.Context, hookId uuid.UUID) error
	CreateHook(ctx context.Context, params *hook.HookCreateParams) (*models.Hook, error)
}

type HookUsecase struct {
	hookSvc HookServiceIF
}

func NewHookUsecase(hookSvc HookServiceIF) *HookUsecase {
	return &HookUsecase{
		hookSvc: hookSvc,
	}
}

func (t *HookUsecase) GetHook(ctx context.Context, hookId uuid.UUID) (*models.Hook, error) {
	_logger := logger.FromContext(ctx).With().Stringer("hookId", hookId).Logger()
	return t.hookSvc.GetHook(logger.WithContext(ctx, _logger), hookId)
}

func (t *HookUsecase) DeleteHook(ctx context.Context, hookId uuid.UUID) error {
	_logger := logger.FromContext(ctx).With().Stringer("hookId", hookId).Logger()
	return t.hookSvc.DeleteHook(logger.WithContext(ctx, _logger), hookId)
}

func (t *HookUsecase) CreateHook(ctx context.Context, params *hook.HookCreateParams) (*models.Hook, error) {
	return t.hookSvc.CreateHook(ctx, params)
}
