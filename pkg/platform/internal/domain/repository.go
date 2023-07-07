package domain

import "mrcAPI/pkg/platform"

type PlatformRepository interface {
	GetAllPlatforms() ([]platform.Platform, error)
}
