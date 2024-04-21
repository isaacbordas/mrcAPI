package application

import (
	"mrcAPI/pkg/game"
	"mrcAPI/pkg/importer/internal/infrastructure"
)

type Importer struct {
	ga infrastructure.GameImporterAPIRepository
	d  infrastructure.DeveloperImporterAPIRepository
	pu infrastructure.PublisherImporterAPIRepository
	g  infrastructure.GenreImporterAPIRepository
	p  infrastructure.PlatformImporterAPIRepository
	r  infrastructure.ImporterMysqlRepository
}

func NewImporterService(
	ga infrastructure.GameImporterAPIRepository,
	d infrastructure.DeveloperImporterAPIRepository,
	pu infrastructure.PublisherImporterAPIRepository,
	g infrastructure.GenreImporterAPIRepository,
	p infrastructure.PlatformImporterAPIRepository,
	r infrastructure.ImporterMysqlRepository,
) Importer {
	return Importer{ga: ga, d: d, pu: pu, g: g, p: p, r: r}
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

func (i Importer) ImportPublishers() error {
	publishersAPI, errAPI := i.pu.GetPublishers()
	if errAPI != nil {
		return errAPI
	}

	return i.r.PersistPublishers(publishersAPI)
}

func (i Importer) ImportGameByName(name string) ([]game.Game, error) {
	games, errAPI := i.ga.GetGameByName(name)
	if errAPI != nil {
		return nil, errAPI
	}

	return games, nil
}
