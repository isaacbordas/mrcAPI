package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mrcAPI/pkg/platform/wire"
	"mrcAPI/pkg/system/config"
	"net/http"
)

func main() {
	host, errHost := config.GoDotEnvVariable("HOST")
	if errHost != nil {
		panic(errHost)
	}

	port, errPort := config.GoDotEnvVariable("PORT")
	if errPort != nil {
		panic(errPort)
	}

	router := gin.Default()
	router.GET("/platforms", getPlatforms)

	errRun := router.Run(fmt.Sprintf("%s:%s", host, port))
	if errRun != nil {
		panic(errRun)
	}
}

func getPlatforms(c *gin.Context) {
	plat, err := wire.ProvidePlatform()
	if err != nil {
		panic(err)
	}

	platforms, errPlat := plat.GetPlatforms()
	if errPlat != nil {
		c.JSON(http.StatusBadRequest, errPlat.Error())
	}

	c.JSON(http.StatusOK, platforms)
}
