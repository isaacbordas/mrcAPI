package importer

type Importer interface {
	ImportPlatforms() error
}
