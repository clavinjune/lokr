package pkg

import (
	"context"
	"github.com/jmoiron/sqlx"
)

var _ TxHandler = (*TxHandlerFunc)(nil)

type TxHandler interface {
	Handle(context.Context, *sqlx.Tx) error
}

type TxHandlerFunc func(context.Context, *sqlx.Tx) error

func (t TxHandlerFunc) Handle(ctx context.Context, tx *sqlx.Tx) error {
	return t(ctx, tx)
}
