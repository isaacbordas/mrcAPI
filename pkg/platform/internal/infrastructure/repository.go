package infrastructure

import (
	"database/sql"
	"mrcAPI/pkg/platform"
)

type PlatformMysqlRepository struct {
	db *sql.DB
}

func NewPlatformMysqlRepository(db *sql.DB) PlatformMysqlRepository {
	return PlatformMysqlRepository{db: db}
}

func (p PlatformMysqlRepository) GetAllPlatforms() ([]platform.Platform, error) {
	var platforms []platform.Platform

	query := "SELECT p.platform_uuid, p.name, p.slug FROM platforms p ORDER BY p.name"

	result, err := p.db.Query(query)
	if err != nil {
		return platforms, err
	}

	var platformUuid, name, slug string
	for result.Next() {
		err = result.Scan(&platformUuid, &name, &slug)
		if err != nil {
			result.Close()
			return platforms, err
		}
		platforms = append(platforms, platform.Platform{
			UUID: platformUuid,
			Name: name,
			Slug: slug,
		})
	}

	return platforms, nil
}
