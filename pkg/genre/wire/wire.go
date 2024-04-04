//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"mrcAPI/pkg/genre"
	"mrcAPI/pkg/genre/internal/application"
	"mrcAPI/pkg/genre/internal/infrastructure"
	"mrcAPI/pkg/system"
)

func provideGenreMysqlRepository() (infrastructure.GenreMysqlRepository, error) {
	panic(wire.Build(
		system.ProvideMySQL,
		infrastructure.NewGenreMysqlRepository,
	))
}

func ProvideGenre() (genre.IGenre, error) {
	panic(wire.Build(
		provideGenreMysqlRepository,
		application.NewGenreService,
		wire.Bind(new(genre.IGenre), new(application.Genre)),
	))
}
