package application

import (
	"mrcAPI/pkg/platform"
	"mrcAPI/pkg/platform/internal/domain"
)

type Platform struct {
	s domain.PlatformService
}

func NewPlatformService(s domain.PlatformService) Platform {
	return Platform{s: s}
}

func (p Platform) GetPlatforms() ([]platform.Platform, error) {
	return p.s.GetPlatforms()
}
