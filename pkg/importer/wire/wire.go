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

func provideDeveloperImporterAPIRepository() (infrastructure.DeveloperImporterAPIRepository, error) {
	panic(wire.Build(
		system.ProvideClient,
		infrastructure.NewDeveloperImporterAPIRepository,
	))
}

func providePublisherImporterAPIRepository() (infrastructure.PublisherImporterAPIRepository, error) {
	panic(wire.Build(
		system.ProvideClient,
		infrastructure.NewPublisherImporterAPIRepository,
	))
}

func provideGameImporterAPIRepository() (infrastructure.GameImporterAPIRepository, error) {
	panic(wire.Build(
		system.ProvideClient,
		infrastructure.NewGameImporterAPIRepository,
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
		provideGameImporterAPIRepository,
		providePublisherImporterAPIRepository,
		provideDeveloperImporterAPIRepository,
		provideGenreImporterAPIRepository,
		providePlatformImporterAPIRepository,
		provideImporterMysqlRepository,
		application.NewImporterService,
		wire.Bind(new(importer.Importer), new(application.Importer)),
	))
}
