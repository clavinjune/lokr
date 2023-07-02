//go:build wireinject
// +build wireinject

package lokr

import (
	"context"
	"database/sql"
	"github.com/google/wire"
)

func Wire(ctx context.Context, db *sql.DB, driverName string, isDebugMode bool) (*Lokr, error) {
	panic(wire.Build(ProviderSet))
}
