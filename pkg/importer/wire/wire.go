//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"mrcAPI/pkg/importer"
	"mrcAPI/pkg/importer/internal/application"
	"mrcAPI/pkg/importer/internal/infrastructure"
	"mrcAPI/pkg/system"
)

func providePlatformImporterAPIRepository() (infrastructure.PlatformImporterAPIRepository, error) {
	panic(wire.Build(
		system.ProvideClient,
		infrastructure.NewPlatformImporterAPIRepository,
	))
}

func provideGenreImporterAPIRepository() (infrastructure.GenreImporterAPIRepository, error) {
	panic(wire.Build(
		system.ProvideClient,
		infrastructure.NewGenreImporterAPIRepository,
	))
}

func provideImporterMysqlRepository() (infrastructure.ImporterMysqlRepository, error) {
	panic(wire.Build(
		system.ProvideMySQL,
		infrastructure.NewImporterMysqlRepository,
	))
}

func ProvideImport() (importer.Importer, error) {
	panic(wire.Build(
		provideGenreImporterAPIRepository,
		providePlatformImporterAPIRepository,
		provideImporterMysqlRepository,
		application.NewImporterService,
		wire.Bind(new(importer.Importer), new(application.Importer)),
	))
}
