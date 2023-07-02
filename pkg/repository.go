package pkg

import (
	"context"

	lokrv1 "github.com/clavinjune/lokr/api/lokr/v1"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	StoreLock(ctx context.Context, lock *lokrv1.Lock) error
	FetchLockTx(ctx context.Context, tx *sqlx.Tx, key string) (*lokrv1.Lock, error)
	PatchLockByKey(ctx context.Context, lock *lokrv1.Lock) error
	PatchLockByKeyTx(ctx context.Context, tx *sqlx.Tx, lock *lokrv1.Lock) error
	Tx(ctx context.Context, fn TxHandler) error
}
