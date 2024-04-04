package infrastructure

import (
	"database/sql"
	"fmt"
	"mrcAPI/pkg/genre"
	"mrcAPI/pkg/platform"
	"strings"
)

type ImporterMysqlRepository struct {
	db *sql.DB
}

type preparedQuery struct {
	query  string
	values []any
}

func NewImporterMysqlRepository(db *sql.DB) ImporterMysqlRepository {
	return ImporterMysqlRepository{db: db}
}

func (i ImporterMysqlRepository) PersistPlatforms(ps []platform.Platform) error {
	query := preparePlatformInsert(ps)
	_, err := i.db.Exec(query.query, query.values...)
	return err
}

func preparePlatformInsert(p []platform.Platform) preparedQuery {
	params := make([]string, 0, len(p))
	values := make([]any, 0, len(p))

	for _, platformItem := range p {
		params = append(params, "(?,?,?,?)")
		values = append(values, platformItem.UUID, platformItem.ApiID, platformItem.Name, platformItem.Slug)
	}

	query := fmt.Sprintf("INSERT INTO platforms(`platform_uuid`,`platform_id`,`platform_name`,`platform_slug`) VALUES %s ON DUPLICATE KEY UPDATE platform_id=VALUES(platform_id)", strings.Join(params, ","))

	return preparedQuery{
		query:  query,
		values: values,
	}
}

func (i ImporterMysqlRepository) PersistGenres(ps []genre.Genre) error {
	query := prepareGenreInsert(ps)
	_, err := i.db.Exec(query.query, query.values...)
	return err
}

func prepareGenreInsert(p []genre.Genre) preparedQuery {
	params := make([]string, 0, len(p))
	values := make([]any, 0, len(p))

	for _, genreItem := range p {
		params = append(params, "(?,?,?)")
		values = append(values, genreItem.UUID, genreItem.ApiID, genreItem.Name)
	}

	query := fmt.Sprintf("INSERT INTO genres(`genre_uuid`,`genre_id`,`genre_name`) VALUES %s ON DUPLICATE KEY UPDATE genre_id=VALUES(genreItem.ApiID)", strings.Join(params, ","))

	return preparedQuery{
		query:  query,
		values: values,
	}
}
