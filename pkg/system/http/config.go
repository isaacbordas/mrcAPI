package http

import "github.com/kelseyhightower/envconfig"

type APIConfig struct {
	Host   string `default:"https://api.thegamesdb.net/" envconfig:"API_HOST"`
	APIKey string `default:"" envconfig:"API_KEY"`
}

func NewAPIConfig() APIConfig {
	cfg := APIConfig{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
