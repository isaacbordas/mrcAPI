package infrastructure

import (
	"database/sql"
	"fmt"
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
	query := prepareInsert(ps)
	_, err := i.db.Exec(query.query, query.values...)
	return err
}

func prepareInsert(p []platform.Platform) preparedQuery {
	params := make([]string, 0, len(p))
	values := make([]any, 0, len(p))

	for _, platformItem := range p {
		params = append(params, "(?,?,?,?)")
		values = append(values, platformItem.UUID, platformItem.ApiID, platformItem.Name, platformItem.Slug)
	}

	query := fmt.Sprintf("INSERT INTO platforms(`platform_uuid`,`platform_id`,`platform_name`,`platform_slug`) VALUES %s", strings.Join(params, ","))

	return preparedQuery{
		query:  query,
		values: values,
	}
}
