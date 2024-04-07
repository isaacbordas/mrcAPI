package routes

import (
	"github.com/gin-gonic/gin"
	"mrcAPI/pkg/publisher/wire"
	"net/http"
)

func GetPublishers(c *gin.Context) {
	pub, err := wire.ProvidePublisher()
	if err != nil {
		panic(err)
	}

	pubres, errPub := pub.GetPublishers()
	if errPub != nil {
		c.JSON(http.StatusBadRequest, errPub.Error())
	}

	c.JSON(http.StatusOK, pubres)
}
