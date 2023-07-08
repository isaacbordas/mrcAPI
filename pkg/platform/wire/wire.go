//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"mrcAPI/pkg/platform"
	"mrcAPI/pkg/platform/internal/application"
	"mrcAPI/pkg/platform/internal/infrastructure"
	"mrcAPI/pkg/system"
)

func providePlatformMysqlRepository() (infrastructure.PlatformMysqlRepository, error) {
	panic(wire.Build(
		system.ProvideMySQL,
		infrastructure.NewPlatformMysqlRepository,
	))
}

func ProvidePlatform() (platform.IPlatform, error) {
	panic(wire.Build(
		providePlatformMysqlRepository,
		application.NewPlatformService,
		wire.Bind(new(platform.IPlatform), new(application.Platform)),
	))
}
