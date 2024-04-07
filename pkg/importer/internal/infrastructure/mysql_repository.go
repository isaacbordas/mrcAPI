package infrastructure

import (
	"database/sql"
	"fmt"
	"mrcAPI/pkg/developer"
	"mrcAPI/pkg/genre"
	"mrcAPI/pkg/platform"
	"mrcAPI/pkg/publisher"
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

	query := fmt.Sprintf("INSERT INTO genres(`genre_uuid`,`genre_id`,`genre_name`) VALUES %s ON DUPLICATE KEY UPDATE genre_id=VALUES(genre_id)", strings.Join(params, ","))

	return preparedQuery{
		query:  query,
		values: values,
	}
}

func (i ImporterMysqlRepository) PersistDevelopers(ps []developer.Developer) error {
	query := prepareDeveloperInsert(ps)
	_, err := i.db.Exec(query.query, query.values...)
	return err
}

func prepareDeveloperInsert(p []developer.Developer) preparedQuery {
	params := make([]string, 0, len(p))
	values := make([]any, 0, len(p))

	for _, developerItem := range p {
		params = append(params, "(?,?,?)")
		values = append(values, developerItem.UUID, developerItem.ApiID, developerItem.Name)
	}

	query := fmt.Sprintf("INSERT INTO developers(`developer_uuid`,`developer_id`,`developer_name`) VALUES %s ON DUPLICATE KEY UPDATE developer_id=VALUES(developer_id)", strings.Join(params, ","))

	return preparedQuery{
		query:  query,
		values: values,
	}
}

func (i ImporterMysqlRepository) PersistPublishers(ps []publisher.Publisher) error {
	query := preparePublisherInsert(ps)
	_, err := i.db.Exec(query.query, query.values...)
	return err
}

func preparePublisherInsert(p []publisher.Publisher) preparedQuery {
	params := make([]string, 0, len(p))
	values := make([]any, 0, len(p))

	for _, publisherItem := range p {
		params = append(params, "(?,?,?)")
		values = append(values, publisherItem.UUID, publisherItem.ApiID, publisherItem.Name)
	}

	query := fmt.Sprintf("INSERT INTO publishers(`publisher_uuid`,`publisher_id`,`publisher_name`) VALUES %s ON DUPLICATE KEY UPDATE publisher_id=VALUES(publisher_id)", strings.Join(params, ","))

	return preparedQuery{
		query:  query,
		values: values,
	}
}
