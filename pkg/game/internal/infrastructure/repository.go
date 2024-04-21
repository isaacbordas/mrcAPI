package infrastructure

import (
	"database/sql"
	"mrcAPI/pkg/errors"
	"mrcAPI/pkg/game"
)

type GameMysqlRepository struct {
	db *sql.DB
}

func NewGameMysqlRepository(db *sql.DB) GameMysqlRepository {
	return GameMysqlRepository{db: db}
}

func (g GameMysqlRepository) GetGameByName(name string) ([]game.Game, error) {
	var games []game.Game

	fullNameQuery := "%" + name + "%"
	query := "SELECT g.game_uuid, g.game_name, g.platform, g.region FROM games g WHERE g.game_name LIKE ?"
	result, err := g.db.Query(query, fullNameQuery)
	if err != nil {
		return games, err
	}

	var gameUuid, gameName string
	var gamePlatform, gameRegion int32
	for result.Next() {
		err = result.Scan(&gameUuid, &gameName, &gamePlatform, &gameRegion)
		if err != nil {
			result.Close()
			return games, err
		}
		games = append(games, game.Game{
			UUID:     gameUuid,
			Name:     gameName,
			Platform: gamePlatform,
			Region:   gameRegion,
		})
	}

	if len(games) < 1 {
		return games, errors.ErrItemNotFound{Entity: "Game", Slug: name}
	}

	return games, nil
}
