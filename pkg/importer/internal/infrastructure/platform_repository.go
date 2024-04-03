package infrastructure

import (
	"encoding/json"
	"github.com/google/uuid"
	"mrcAPI/pkg/errors"
	"mrcAPI/pkg/platform"
	"mrcAPI/pkg/system/http"
)

const platformEndpoint = "v1/Platforms"

type PlatformImporterRepository interface {
	GetPlatforms() ([]platform.Platform, error)
}

type PlatformImporterAPIRepository struct {
	client http.Client
}

type PlatformsResponseRoot struct {
	Code   int                   `json:"code"`
	Status string                `json:"status"`
	Data   PlatformsResponseData `json:"data"`
}

type PlatformsResponseData struct {
	Count     int32                `json:"count"`
	Platforms map[string]Platforms `json:"platforms"`
}

type Platforms struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"alias"`
}

func NewPlatformImporterAPIRepository(client http.Client) PlatformImporterAPIRepository {
	return PlatformImporterAPIRepository{client: client}
}

func (i PlatformImporterAPIRepository) GetPlatforms() ([]platform.Platform, error) {
	var platforms []platform.Platform

	response, err := i.client.GetRequest(platformEndpoint, nil)
	if err != nil {
		return nil, err
	}

	transformed, errTransform := i.transformResponse(response)
	if errTransform != nil {
		return nil, errTransform
	}

	for _, plat := range transformed.Platforms {
		platforms = append(platforms, platform.Platform{
			UUID:  uuid.NewString(),
			ApiID: plat.ID,
			Name:  plat.Name,
			Slug:  plat.Slug,
		})
	}

	return platforms, err
}

func (i PlatformImporterAPIRepository) transformResponse(body []byte) (PlatformsResponseData, error) {
	platformsResponseRoot := PlatformsResponseRoot{}
	err := json.Unmarshal(body, &platformsResponseRoot)
	if err != nil {
		return PlatformsResponseData{}, err
	}

	if platformsResponseRoot.Data.Count < 1 {
		return PlatformsResponseData{}, errors.ErrPlatformsNotFoundAPI{}
	}

	return platformsResponseRoot.Data, nil
}
