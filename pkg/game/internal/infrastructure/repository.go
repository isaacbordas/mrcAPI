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

func (g GameMysqlRepository) GetAllGames() ([]game.Game, error) {
	var games []game.Game

	query := "SELECT g.game_uuid, g.game_name, g.game_slug FROM games g ORDER BY g.game_name"

	result, err := g.db.Query(query)
	if err != nil {
		return games, err
	}

	var gameUuid, gameName, gameSlug string
	for result.Next() {
		err = result.Scan(&gameUuid, &gameName, &gameSlug)
		if err != nil {
			result.Close()
			return games, err
		}
		games = append(games, game.Game{
			UUID: gameUuid,
			Name: gameName,
			Slug: gameSlug,
		})
	}

	return games, nil
}

func (g GameMysqlRepository) GetGameBySlug(slug string) (game.Game, error) {
	var gameItem game.Game

	query := "SELECT p.game_uuid, p.game_name, p.game_slug FROM games p WHERE game_slug=?"
	result := g.db.QueryRow(query, slug)

	var gameUuid, gameName, gameSlug string
	err := result.Scan(&gameUuid, &gameName, &gameSlug)
	switch err {
	case sql.ErrNoRows:
		return gameItem, errors.ErrItemNotFound{Entity: "Game", Slug: slug}
	case nil:
		return game.Game{
			UUID: gameUuid,
			Name: gameName,
			Slug: gameSlug,
		}, nil
	default:
		return gameItem, err
	}
}
