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
	platforms := []platform.Platform{
		{Name: "MSX", Slug: "msx"},
		{Name: "Neo Geo", Slug: "neo-geo"},
	}
	return platforms, nil
}
