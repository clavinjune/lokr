package lokr

import (
	"database/sql"

	"github.com/clavinjune/lokr/pkg"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var (
	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideLokr,
		ProvideLokrRepository,
		wire.Bind(new(pkg.Repository), new(*repository)),
	)
)

func ProvideLokr(repo pkg.Repository, isDebugMode bool) (*Lokr, error) {
	return &Lokr{
		repo:        repo,
		isDebugMode: isDebugMode,
	}, nil
}

func ProvideLokrRepository(db *sql.DB, driverName string) (*repository, error) {
	if db == nil {
		return nil, pkg.ErrNilDB
	}

	dbx := sqlx.NewDb(db, driverName)
	dbx.Mapper = reflectx.NewMapper("json")
	return &repository{
		db: dbx,
	}, nil
}
