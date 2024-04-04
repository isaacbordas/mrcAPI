package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"mrcAPI/cmd/app/routes"
	"mrcAPI/cmd/app/routes/admin"
)

type ServerConfig struct {
	Host string `default:"0.0.0.0" envconfig:"HOST"`
	Port string `default:"9000" envconfig:"PORT"`
}

func NewServerConfig() ServerConfig {
	cfg := ServerConfig{}
	envconfig.MustProcess("", &cfg)
	return cfg
}

func main() {
	cfg := NewServerConfig()

	router := gin.Default()
	router.GET("/platforms", routes.GetPlatforms)
	router.GET("/platform/:slug", routes.GetPlatformBySlug)
	router.GET("/genres", routes.GetGenres)
	router.GET("/games", routes.GetGames)
	router.GET("/game/:slug", routes.GetGameBySlug)
	router.GET("/admin/importer/:entity", admin.ImportByEntity)

	errRun := router.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if errRun != nil {
		panic(errRun)
	}
}
