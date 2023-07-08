package routes

import (
	"github.com/gin-gonic/gin"
	"mrcAPI/pkg/platform/wire"
	"net/http"
)

func GetPlatforms(c *gin.Context) {
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

func GetPlatformBySlug(c *gin.Context) {
	plat, err := wire.ProvidePlatform()
	if err != nil {
		panic(err)
	}

	slug := c.Param("slug")
	platform, errPlat := plat.GetPlatform(slug)
	if errPlat != nil {
		c.JSON(http.StatusBadRequest, errPlat.Error())
	}

	c.JSON(http.StatusOK, platform)
}
