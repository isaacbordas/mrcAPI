package infrastructure

import (
	"encoding/json"
	"mrcAPI/pkg/errors"
	"mrcAPI/pkg/platform"
	"mrcAPI/pkg/system/http"
)

const platformEndpoint = "v1/Platforms"

type ImporterAPIRepository struct {
	client http.Client
}

type PlatformsResponseRoot struct {
	Data PlatformsResponseData `json:"data"`
}

type PlatformsResponseData struct {
	Count     int32       `json:"count"`
	Platforms []Platforms `json:"platforms"`
}

type Platforms struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"alias"`
}

func NewImporterAPIRepository(client http.Client) ImporterAPIRepository {
	return ImporterAPIRepository{client: client}
}

func (i ImporterAPIRepository) GetPlatforms() ([]platform.Platform, error) {
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
			ApiID: plat.ID,
			Name:  plat.Name,
			Slug:  plat.Slug,
		})
	}

	return platforms, err
}

func (i ImporterAPIRepository) transformResponse(body string) (PlatformsResponseData, error) {
	platformsResponseRoot := PlatformsResponseRoot{}
	err := json.Unmarshal([]byte(body), &platformsResponseRoot)
	if err != nil {
		return PlatformsResponseData{}, err
	}

	if len(platformsResponseRoot.Data.Platforms) < 1 {
		return PlatformsResponseData{}, errors.ErrPlatformsNotFoundAPI{}
	}

	return platformsResponseRoot.Data, nil
}
