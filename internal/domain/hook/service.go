package hook

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/pei223/hook-scheduler/pkg/db"
	"resty.dev/v3"
)

type HookRepo interface {
	GetHook(ctx context.Context, tx *sql.Tx, hookId uuid.UUID) (*models.Hook, error)
	DeleteHook(ctx context.Context, tx *sql.Tx, hookId uuid.UUID) error
	CreateHook(ctx context.Context, tx *sql.Tx, params *HookCreateParams) (*models.Hook, error)
	GetAllHooks(ctx context.Context, tx *sql.Tx, limit int, offset int) (models.HookSlice, error)
}

type HookService struct {
	db        *sql.DB
	repo      HookRepo
	apiClient *resty.Client
}

func NewHookService(db *sql.DB, repo HookRepo, apiClient *resty.Client) *HookService {
	return &HookService{
		repo:      repo,
		db:        db,
		apiClient: apiClient,
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

func (h *HookService) GetAllHooks(ctx context.Context, limit, offset int) (models.HookSlice, error) {
	var hooks models.HookSlice
	err := db.ReadOnlyTx(
		ctx,
		h.db,
		func(ctx context.Context, tx *sql.Tx) error {
			var err error
			hooks, err = h.repo.GetAllHooks(ctx, tx, limit, offset)
			return err
		},
		nil,
	)
	return hooks, err
}

func (h *HookService) ExecHookInTx(ctx context.Context, tx *sql.Tx, hookID uuid.UUID) (int, error) {
	// TODO リトライ処理
	hook, err := h.repo.GetHook(ctx, tx, hookID)
	if err != nil {
		return 0, err
	}

	headers := map[string]string{}
	for k, v := range hook.Headers {
		headers[k] = fmt.Sprintf("%v", v)
	}
	headers["Content-Type"] = "application/json"

	res, err := h.apiClient.R().SetHeaders(
		headers,
	).SetBody(hook.Body).Execute(hook.Method, hook.URL)
	return res.StatusCode(), err
}
