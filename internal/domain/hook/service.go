package hook

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/db"
)

type HookService struct {
	db   *sql.DB
	repo HookRepo
}

func NewHookService(db *sql.DB, repo HookRepo) *HookService {
	return &HookService{
		repo: repo,
		db:   db,
	}
}

func (h *HookService) GetHook(ctx context.Context, hookId uuid.UUID) (*models.Hook, error) {
	var hook *models.Hook
	err := db.ReadOnlyTx(
		ctx,
		h.db,
		func(ctx context.Context, tx *sql.Tx) error {
			var err error
			hook, err = h.GetHookInTx(ctx, tx, hookId)
			return err
		},
		nil,
	)
	return hook, err
}

func (h *HookService) GetHookInTx(ctx context.Context, tx *sql.Tx, hookId uuid.UUID) (*models.Hook, error) {
	return h.repo.GetHook(ctx, tx, hookId)
}

func (h *HookService) DeleteHook(ctx context.Context, hookId uuid.UUID) error {
	err := db.ExecTx(
		ctx,
		h.db,
		func(ctx context.Context, tx *sql.Tx) error {
			return h.repo.DeleteHook(ctx, tx, hookId)
		},
		nil,
	)
	return err
}

func (h *HookService) DeleteHookInTx(ctx context.Context, tx *sql.Tx, hookId uuid.UUID) error {
	return h.repo.DeleteHook(ctx, tx, hookId)
}

func (h *HookService) CreateHook(ctx context.Context, params *HookCreateParams) (*models.Hook, error) {
	var hook *models.Hook
	err := db.ExecTx(
		ctx,
		h.db,
		func(ctx context.Context, tx *sql.Tx) error {
			var err error
			hook, err = h.CreateHookInTx(ctx, tx, params)
			return err
		},
		nil,
	)
	return hook, err
}

func (h *HookService) CreateHookInTx(ctx context.Context, tx *sql.Tx, params *HookCreateParams) (*models.Hook, error) {
	return h.repo.CreateHook(ctx, tx, params)
}
