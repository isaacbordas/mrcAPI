package application

import (
	"mrcAPI/pkg/platform"
	"mrcAPI/pkg/platform/internal/infrastructure"
)

type Platform struct {
	r infrastructure.PlatformMysqlRepository
}

func NewPlatformService(r infrastructure.PlatformMysqlRepository) Platform {
	return Platform{r: r}
}

func (p Platform) GetPlatforms() ([]platform.Platform, error) {
	return p.r.GetAllPlatforms()
}
