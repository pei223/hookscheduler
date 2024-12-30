package db

import (
	"context"
	"database/sql"
)

func ExecTx(
	ctx context.Context,
	db *sql.DB,
	f func(ctx context.Context, tx *sql.Tx) error,
) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := f(ctx, tx); err != nil {
		return err
	}

	return tx.Commit()
}
