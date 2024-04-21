package infrastructure

import (
	"encoding/json"
	"mrcAPI/pkg/errors"
	"mrcAPI/pkg/game"
	"mrcAPI/pkg/system/http"
)

const gameEndpoint = "v1.1/Games/ByGameName"

type GameImporterRepository interface {
	GetGameByName(gameName string) ([]game.Game, error)
}

type GameImporterAPIRepository struct {
	client http.Client
}

type GamesResponseRoot struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   GamesResponseData `json:"data"`
}

type GamesResponseData struct {
	Count int32   `json:"count"`
	Games []Games `json:"games"`
}

type Games struct {
	ID          int32  `json:"id"`
	Name        string `json:"game_title"`
	Platform    int32  `json:"platform"`
	Region      int32  `json:"region_id"`
	ReleaseDate string `json:"release_date"`
}

func NewGameImporterAPIRepository(client http.Client) GameImporterAPIRepository {
	return GameImporterAPIRepository{client: client}
}

func (i GameImporterAPIRepository) GetGameByName(gameName string) ([]game.Game, error) {
	var games []game.Game

	params := map[string]string{"name": gameName}
	response, err := i.client.GetRequest(gameEndpoint, params)
	if err != nil {
		return nil, err
	}

	transformed, errTransform := i.transformResponse(response, gameName)
	if errTransform != nil {
		return nil, errTransform
	}

	for _, g := range transformed.Games {
		games = append(games, game.Game{
			Name:     g.Name,
			Platform: g.Platform,
			Region:   g.Region,
		})
	}

	return games, err
}

func (i GameImporterAPIRepository) transformResponse(body []byte, gameName string) (GamesResponseData, error) {
	gamesResponseRoot := GamesResponseRoot{}
	err := json.Unmarshal(body, &gamesResponseRoot)
	if err != nil {
		return GamesResponseData{}, err
	}

	if gamesResponseRoot.Data.Count < 1 {
		return GamesResponseData{}, errors.ErrGamesNotFoundAPI{}.Error(gameName)
	}

	return gamesResponseRoot.Data, nil
}
