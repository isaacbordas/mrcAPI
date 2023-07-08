package routes

import (
	"github.com/gin-gonic/gin"
	"mrcAPI/pkg/game/wire"
	"net/http"
)

func GetGames(c *gin.Context) {
	game, err := wire.ProvideGame()
	if err != nil {
		panic(err)
	}

	games, errGames := game.GetGames()
	if errGames != nil {
		c.JSON(http.StatusBadRequest, errGames.Error())
	}

	c.JSON(http.StatusOK, games)
}
