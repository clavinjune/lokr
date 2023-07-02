package lokr

import (
	"context"

	lokrv1 "github.com/clavinjune/lokr/api/lokr/v1"
	"github.com/clavinjune/lokr/pkg"
	"github.com/jmoiron/sqlx"
)

var (
	_ pkg.Repository = (*repository)(nil)
)

type repository struct {
	db *sqlx.DB
}

func (r *repository) StoreLock(ctx context.Context, lock *lokrv1.Lock) error {
	result, err := r.db.NamedExecContext(ctx, `INSERT INTO locks (key, is_locked) VALUES(:key, :is_locked)`, lock)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return pkg.ErrZeroAffected
	}

	return nil
}

func (r *repository) FetchLockTx(ctx context.Context, tx *sqlx.Tx, key string) (*lokrv1.Lock, error) {
	rows, err := sqlx.NamedQueryContext(ctx, tx, `SELECT * FROM locks WHERE key = :key FOR UPDATE`, map[string]any{
		"key": key,
	})
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()

	var lock lokrv1.Lock
	if rows.Next() {
		if err := rows.StructScan(&lock); err != nil {
			return nil, err
		}
	}

	return &lock, nil
}

func (r *repository) PatchLockByKey(ctx context.Context, lock *lokrv1.Lock) error {
	return r.patchLockByKey(ctx, r.db, lock)
}
func (r *repository) PatchLockByKeyTx(ctx context.Context, tx *sqlx.Tx, lock *lokrv1.Lock) error {
	return r.patchLockByKey(ctx, tx, lock)
}

func (r *repository) Tx(ctx context.Context, fn pkg.TxHandler) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	if err := fn.Handle(ctx, tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *repository) patchLockByKey(ctx context.Context, e sqlx.ExtContext, lock *lokrv1.Lock) error {
	result, err := sqlx.NamedExecContext(ctx, e, `UPDATE locks SET is_locked = :is_locked WHERE key = :key`, lock)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n < 1 {
		return pkg.ErrZeroAffected
	}

	return nil
}
