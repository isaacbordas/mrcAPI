package infrastructure

import (
	"database/sql"
	"mrcAPI/pkg/errors"
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

	query := "SELECT p.platform_uuid, p.platform_id, p.platform_name, p.platform_slug FROM platforms p ORDER BY p.platform_name"

	result, err := p.db.Query(query)
	if err != nil {
		return platforms, err
	}

	var platformUuid, platformName, platformSlug string
	var platformID int32
	for result.Next() {
		err = result.Scan(&platformUuid, &platformID, &platformName, &platformSlug)
		if err != nil {
			result.Close()
			return platforms, err
		}
		platforms = append(platforms, platform.Platform{
			UUID:  platformUuid,
			ApiID: platformID,
			Name:  platformName,
			Slug:  platformSlug,
		})
	}

	return platforms, nil
}

func (p PlatformMysqlRepository) GetPlatformBySlug(slug string) (platform.Platform, error) {
	var plat platform.Platform

	query := "SELECT p.platform_uuid, p.platform_id, p.platform_name, p.platform_slug FROM platforms p WHERE platform_slug=?"
	result := p.db.QueryRow(query, slug)

	var platformUuid, platformName, platformSlug string
	var platformID int32
	err := result.Scan(&platformUuid, &platformID, &platformName, &platformSlug)
	switch err {
	case sql.ErrNoRows:
		return plat, errors.ErrItemNotFound{Entity: "Platform", Slug: slug}
	case nil:
		return platform.Platform{
			UUID:  platformUuid,
			ApiID: platformID,
			Name:  platformName,
			Slug:  platformSlug,
		}, nil
	default:
		return plat, err
	}
}
