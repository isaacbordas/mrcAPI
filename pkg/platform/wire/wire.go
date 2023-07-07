//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"mrcAPI/pkg/platform"
	"mrcAPI/pkg/platform/internal"
)

func ProvidePlatform() (platform.IPlatform, error) {
	panic(wire.Bind(
		new(platform.IPlatform), new(internal.Platform),
	))
}
