package hook

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/pei223/hook-scheduler/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type HookRepo interface {
	GetHook(ctx context.Context, tx *sql.Tx, hookId uuid.UUID) (*models.Hook, error)
	DeleteHook(ctx context.Context, tx *sql.Tx, hookId uuid.UUID) error
	CreateHook(ctx context.Context, tx *sql.Tx, params *HookCreateParams) (*models.Hook, error)
}

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
