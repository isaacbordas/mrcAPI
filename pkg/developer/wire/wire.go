//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"mrcAPI/pkg/developer"
	"mrcAPI/pkg/developer/internal/application"
	"mrcAPI/pkg/developer/internal/infrastructure"
	"mrcAPI/pkg/system"
)

func provideDeveloperMysqlRepository() (infrastructure.DeveloperMysqlRepository, error) {
	panic(wire.Build(
		system.ProvideMySQL,
		infrastructure.NewDeveloperMysqlRepository,
	))
}

func ProvideDeveloper() (developer.IDeveloper, error) {
	panic(wire.Build(
		provideDeveloperMysqlRepository,
		application.NewDeveloperService,
		wire.Bind(new(developer.IDeveloper), new(application.Developer)),
	))
}
