package infrastructure

import (
	"encoding/json"
	"github.com/google/uuid"
	"mrcAPI/pkg/errors"
	"mrcAPI/pkg/genre"
	"mrcAPI/pkg/system/http"
)

const genreEndpoint = "v1/Genres"

type GenreImporterRepository interface {
	GetGenres() ([]genre.Genre, error)
}

type GenreImporterAPIRepository struct {
	client http.Client
}

type GenresResponseRoot struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Data   GenresResponseData `json:"data"`
}

type GenresResponseData struct {
	Count  int32             `json:"count"`
	Genres map[string]Genres `json:"genres"`
}

type Genres struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func NewGenreImporterAPIRepository(client http.Client) GenreImporterAPIRepository {
	return GenreImporterAPIRepository{client: client}
}

func (i GenreImporterAPIRepository) GetGenres() ([]genre.Genre, error) {
	var genres []genre.Genre

	response, err := i.client.GetRequest(genreEndpoint, nil)
	if err != nil {
		return nil, err
	}

	transformed, errTransform := i.transformResponse(response)
	if errTransform != nil {
		return nil, errTransform
	}

	for _, gen := range transformed.Genres {
		genres = append(genres, genre.Genre{
			UUID:  uuid.NewString(),
			ApiID: gen.ID,
			Name:  gen.Name,
		})
	}

	return genres, err
}

func (i GenreImporterAPIRepository) transformResponse(body []byte) (GenresResponseData, error) {
	genresResponseRoot := GenresResponseRoot{}
	err := json.Unmarshal(body, &genresResponseRoot)
	if err != nil {
		return GenresResponseData{}, err
	}

	if genresResponseRoot.Data.Count < 1 {
		return GenresResponseData{}, errors.ErrGenresNotFoundAPI{}
	}

	return genresResponseRoot.Data, nil
}
