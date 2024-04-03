package application

import (
	"mrcAPI/pkg/importer/internal/infrastructure"
)

type Importer struct {
	g infrastructure.GenreImporterAPIRepository
	p infrastructure.PlatformImporterAPIRepository
	r infrastructure.ImporterMysqlRepository
}

func NewImporterService(
	g infrastructure.GenreImporterAPIRepository,
	p infrastructure.PlatformImporterAPIRepository,
	r infrastructure.ImporterMysqlRepository,
) Importer {
	return Importer{g: g, p: p, r: r}
}

func (i Importer) ImportPlatforms() error {
	platformsAPI, errAPI := i.p.GetPlatforms()
	if errAPI != nil {
		return errAPI
	}

	return i.r.PersistPlatforms(platformsAPI)
}

func (i Importer) ImportGenres() error {
	genresAPI, errAPI := i.g.GetGenres()
	if errAPI != nil {
		return errAPI
	}

	return i.r.PersistGenres(genresAPI)
}
