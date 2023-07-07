package internal

import (
	"mrcAPI/pkg/platform"
)

type Platform struct {
}

func (p Platform) GetPlatforms() ([]platform.Platform, error) {
	var platforms = []platform.Platform{
		{Name: "SNES", Slug: "snes"},
		{Name: "MSX", Slug: "msx"},
	}

	return platforms, nil
}
