package importer

import "mrcAPI/pkg/game"

type Importer interface {
	ImportPlatforms() error
	ImportGenres() error
	ImportDevelopers() error
	ImportPublishers() error
	ImportGameByName(name string) ([]game.Game, error)
}
