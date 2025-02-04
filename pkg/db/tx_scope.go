package db

import (
	"context"
	"database/sql"

	"github.com/pei223/hook-scheduler/pkg/logger"
)

func ReadOnlyTx(
	ctx context.Context,
	db *sql.DB,
	f func(ctx context.Context, tx *sql.Tx) error,
	option *sql.TxOptions,
) error {
	newOption := &sql.TxOptions{
		ReadOnly: true,
	}
	if option != nil {
		newOption.Isolation = option.Isolation
	}
	return ExecTx(ctx, db, f, newOption)
}

func ExecTx(
	ctx context.Context,
	db *sql.DB,
	f func(ctx context.Context, tx *sql.Tx) error,
	option *sql.TxOptions,
) error {
	logger := logger.FromContext(ctx)
	var err error
	tx, err := db.BeginTx(ctx, option)
	if err != nil {
		logger.Warn().Err(err).Msg("failed to begin transaction")
		return err
	}
	defer func() {
		if err != nil {
			logger.Warn().Err(err).Msg("transaction failed. try rollback")
			err := tx.Rollback()
			if err != nil {
				logger.Error().Err(err).Msg("failed to rollback")
			}
		}
	}()

	if err = f(ctx, tx); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		logger.Warn().Err(err).Msg("failed to commit")
	}
	return err
}
