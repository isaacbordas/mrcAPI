package routes

import (
	"github.com/gin-gonic/gin"
	"mrcAPI/pkg/developer/wire"
	"net/http"
)

func GetDevelopers(c *gin.Context) {
	dev, err := wire.ProvideDeveloper()
	if err != nil {
		panic(err)
	}

	devres, errDev := dev.GetDevelopers()
	if errDev != nil {
		c.JSON(http.StatusBadRequest, errDev.Error())
	}

	c.JSON(http.StatusOK, devres)
}
