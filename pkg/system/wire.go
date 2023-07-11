//go:build wireinject

package system

import (
	"database/sql"
	"github.com/google/wire"
	"mrcAPI/pkg/system/http"
	"mrcAPI/pkg/system/mysql"
)

func ProvideMySQL() (*sql.DB, error) {
	panic(wire.Build(
		mysql.NewDBConfig,
		mysql.NewDB,
	))
}

func ProvideClient() http.Client {
	panic(wire.Build(
		http.NewAPIConfig,
		http.NewClient,
	))
}
