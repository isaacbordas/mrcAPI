package mysql

import (
	"database/sql"
	"fmt"
	"net/url"
	"reflect"

	"github.com/go-sql-driver/mysql"
)

func NewDB(cfg DBConfig) (*sql.DB, error) {
	driverName := fmt.Sprintf("%s-%s", "mysql", cfg.DBName)
	drivers := sql.Drivers()
	if !isDriverRegistered(drivers, driverName) {
		sql.Register(driverName, &mysql.MySQLDriver{})
	}
	return sql.Open(driverName, dBConnInfo(cfg))
}

// dbConnInfo returns a MySQL connection string for the given host.
func dBConnInfo(cfg DBConfig) string {
	// For the connection string syntax, see
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name

	dsnConfig := url.Values{}
	// interpolateParams=true makes the use of prepared statements explicit.
	// the driver will escape the variables for any requests and send ready-for-use queries to the server.
	// see: https://www.percona.com/blog/2020/06/11/mysql-escaping-on-the-client-side-with-go/
	dsnConfig.Set("interpolateParams", "true")
	// parseTime=true allows parsing from sql results into time.Time.
	dsnConfig.Set("parseTime", "true")
	// multiStatements allows to send multiple statements in each Exec or Query
	dsnConfig.Set("multiStatements", "true")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, dsnConfig.Encode())
}

func isDriverRegistered(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
