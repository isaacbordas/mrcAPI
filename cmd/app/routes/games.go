package routes

import (
	"github.com/gin-gonic/gin"
	"mrcAPI/pkg/game/wire"
	"net/http"
	"strings"
)

func GetGameByName(c *gin.Context) {
	game, err := wire.ProvideGame()
	if err != nil {
		panic(err)
	}

	name := c.Param("name")
	if len(strings.TrimSpace(name)) < 1 {
		c.JSON(http.StatusBadRequest, "game name required")
	}

	games, errGames := game.GetGameByName(name)
	if errGames != nil {
		c.JSON(http.StatusBadRequest, errGames.Error())
	}

	c.JSON(http.StatusOK, games)
}
