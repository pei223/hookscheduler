package hook

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type hookRepo struct {
}

func NewHookRepo() HookRepo {
	return &hookRepo{}
}

func (t *hookRepo) GetHook(ctx context.Context, tx *sql.Tx, hookId uuid.UUID) (*models.Hook, error) {
	return models.Hooks(models.HookWhere.HookID.EQ(hookId)).One(ctx, tx)
}

func (t *hookRepo) DeleteHook(ctx context.Context, tx *sql.Tx, hookId uuid.UUID) error {
	hook, err := models.Hooks(models.HookWhere.HookID.EQ(hookId)).One(ctx, tx)
	if err != nil {
		return err
	}
	_, err = hook.Delete(ctx, tx)
	return err
}

func (t *hookRepo) CreateHook(ctx context.Context, tx *sql.Tx, params *HookCreateParams) (*models.Hook, error) {
	hook := &models.Hook{
		HookID:  uuid.New(),
		URL:     params.URL,
		Method:  params.Method,
		Body:    params.Body,
		Headers: params.Headers,
	}
	err := hook.Insert(ctx, tx, boil.Infer())
	return hook, err
}

func (t *hookRepo) GetAllHooks(ctx context.Context, tx *sql.Tx, limit int, offset int) (models.HookSlice, int, error) {
	total, err := models.Hooks().Count(ctx, tx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get total hooks: %w", err)
	}
	hooks, err := models.Hooks(
		qm.Limit(limit),
		qm.Offset(offset),
		qm.OrderBy(models.HookColumns.DisplayName),
	).All(ctx, tx)
	if err != nil {
		return nil, 0, fmt.Errorf("フックの取得に失敗しました: %w", err)
	}
	return hooks, int(total), err
}
