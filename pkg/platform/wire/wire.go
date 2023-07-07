//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"mrcAPI/pkg/platform"
	"mrcAPI/pkg/platform/internal/application"
	"mrcAPI/pkg/platform/internal/domain"
	"mrcAPI/pkg/platform/internal/infrastructure"
	"mrcAPI/pkg/system"
)

var mysqlRepository = wire.NewSet(infrastructure.NewPlatformMysqlRepository, wire.Bind(new(domain.PlatformRepository), new(infrastructure.PlatformMysqlRepository)))

func providePlatformMysqlRepository() (domain.PlatformRepository, error) {
	panic(wire.Build(
		system.ProvideMySQL,
		mysqlRepository,
	))
}

func ProvidePlatform() (platform.IPlatform, error) {
	panic(wire.Build(
		providePlatformMysqlRepository,
		domain.NewPlatformService,
		application.NewPlatformService,
		wire.Bind(new(platform.IPlatform), new(application.Platform)),
	))
}
