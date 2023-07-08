package application

import (
	"mrcAPI/pkg/game"
	"mrcAPI/pkg/game/internal/infrastructure"
)

type Game struct {
	r infrastructure.GameMysqlRepository
}

func NewGameService(r infrastructure.GameMysqlRepository) Game {
	return Game{r: r}
}

func (g Game) GetGames() ([]game.Game, error) {
	return g.r.GetAllGames()
}
