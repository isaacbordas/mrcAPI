package domain

import "mrcAPI/pkg/platform"

type platformService struct {
	r PlatformRepository
}

type PlatformService interface {
	GetPlatforms() ([]platform.Platform, error)
}

func NewPlatformService(r PlatformRepository) PlatformService {
	return platformService{r: r}
}

func (p platformService) GetPlatforms() ([]platform.Platform, error) {
	return p.r.GetAllPlatforms()
}
