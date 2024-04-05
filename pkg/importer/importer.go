package importer

type Importer interface {
	ImportPlatforms() error
	ImportGenres() error
	ImportDevelopers() error
}
