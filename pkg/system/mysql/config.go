package mysql

import "github.com/kelseyhightower/envconfig"

type DBConfig struct {
	Host     string `default:"localhost" envconfig:"MYSQL_HOST"`
	Port     uint16 `default:"3306" envconfig:"MYSQL_PORT"`
	DBName   string `default:"myretrocollectiondb" envconfig:"MYSQL_DBNAME"`
	User     string `default:"myretrouser" envconfig:"MYSQL_USER"`
	Password string `default:"root" envconfig:"MYSQL_PASSWORD"`
}

func NewDBConfig() DBConfig {
	cfg := DBConfig{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
