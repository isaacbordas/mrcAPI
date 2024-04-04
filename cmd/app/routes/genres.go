package routes

import (
	"github.com/gin-gonic/gin"
	"mrcAPI/pkg/genre/wire"
	"net/http"
)

func GetGenres(c *gin.Context) {
	gen, err := wire.ProvideGenre()
	if err != nil {
		panic(err)
	}

	genres, errGen := gen.GetGenres()
	if errGen != nil {
		c.JSON(http.StatusBadRequest, errGen.Error())
	}

	c.JSON(http.StatusOK, genres)
}
