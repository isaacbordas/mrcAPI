package system

import (
	"database/sql"
	"github.com/google/wire"
	"mrcAPI/pkg/system/mysql"
)

func ProvideMySQL() (*sql.DB, error) {
	panic(wire.Build(
		mysql.NewDBConfig,
		mysql.NewDB,
	))
}
