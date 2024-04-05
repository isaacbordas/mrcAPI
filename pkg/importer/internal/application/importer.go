package application

import (
	"mrcAPI/pkg/importer/internal/infrastructure"
)

type Importer struct {
	d infrastructure.DeveloperImporterAPIRepository
	g infrastructure.GenreImporterAPIRepository
	p infrastructure.PlatformImporterAPIRepository
	r infrastructure.ImporterMysqlRepository
}

func NewImporterService(
	d infrastructure.DeveloperImporterAPIRepository,
	g infrastructure.GenreImporterAPIRepository,
	p infrastructure.PlatformImporterAPIRepository,
	r infrastructure.ImporterMysqlRepository,
) Importer {
	return Importer{d: d, g: g, p: p, r: r}
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

func (i Importer) ImportDevelopers() error {
	developersAPI, errAPI := i.d.GetDevelopers()
	if errAPI != nil {
		return errAPI
	}

	return i.r.PersistDevelopers(developersAPI)
}
