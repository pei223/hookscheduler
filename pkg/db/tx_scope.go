package db

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"
)

func ExecTx(
	ctx context.Context,
	logger *zerolog.Logger,
	db *sql.DB,
	f func(ctx context.Context, tx *sql.Tx) error,
	options ...sql.TxOptions,
) error {
	option := &sql.TxOptions{}
	for _, opt := range options {
		option = &opt
	}

	tx, err := db.BeginTx(ctx, option)
	if err != nil {
		logger.Error().Err(err).Msg("failed to begin transaction")
		return err
	}
	defer func() {
		err := tx.Rollback()
		if err != nil {
			logger.Error().Err(err).Msg("failed to rollback")
		}
	}()

	if err := f(ctx, tx); err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		logger.Error().Err(err).Msg("failed to commit")
	}
	return err
}
