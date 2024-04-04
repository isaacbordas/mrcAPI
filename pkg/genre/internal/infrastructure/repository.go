package infrastructure

import (
	"database/sql"
	"mrcAPI/pkg/genre"
)

type GenreMysqlRepository struct {
	db *sql.DB
}

func NewGenreMysqlRepository(db *sql.DB) GenreMysqlRepository {
	return GenreMysqlRepository{db: db}
}

func (g GenreMysqlRepository) GetAllGenres() ([]genre.Genre, error) {
	var genres []genre.Genre

	query := "SELECT g.genre_uuid, g.genre_id, g.genre_name FROM genres g ORDER BY g.genre_name"

	result, err := g.db.Query(query)
	if err != nil {
		return genres, err
	}

	var genreUuid, genreName string
	var genreID int32
	for result.Next() {
		err = result.Scan(&genreUuid, &genreID, &genreName)
		if err != nil {
			result.Close()
			return genres, err
		}
		genres = append(genres, genre.Genre{
			UUID:  genreUuid,
			ApiID: genreID,
			Name:  genreName,
		})
	}

	return genres, nil
}
