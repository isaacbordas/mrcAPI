package infrastructure

import (
	"encoding/json"
	"github.com/google/uuid"
	"mrcAPI/pkg/developer"
	"mrcAPI/pkg/errors"
	"mrcAPI/pkg/system/http"
)

const developerEndpoint = "v1/Developers"

type DeveloperImporterRepository interface {
	GetDevelopers() ([]developer.Developer, error)
}

type DeveloperImporterAPIRepository struct {
	client http.Client
}

type DevelopersResponseRoot struct {
	Code   int                    `json:"code"`
	Status string                 `json:"status"`
	Data   DevelopersResponseData `json:"data"`
}

type DevelopersResponseData struct {
	Count      int32                 `json:"count"`
	Developers map[string]Developers `json:"developers"`
}

type Developers struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func NewDeveloperImporterAPIRepository(client http.Client) DeveloperImporterAPIRepository {
	return DeveloperImporterAPIRepository{client: client}
}

func (i DeveloperImporterAPIRepository) GetDevelopers() ([]developer.Developer, error) {
	var developers []developer.Developer

	response, err := i.client.GetRequest(developerEndpoint, nil)
	if err != nil {
		return nil, err
	}

	transformed, errTransform := i.transformResponse(response)
	if errTransform != nil {
		return nil, errTransform
	}

	for _, dev := range transformed.Developers {
		developers = append(developers, developer.Developer{
			UUID:  uuid.NewString(),
			ApiID: dev.ID,
			Name:  dev.Name,
		})
	}

	return developers, err
}

func (i DeveloperImporterAPIRepository) transformResponse(body []byte) (DevelopersResponseData, error) {
	developersResponseRoot := DevelopersResponseRoot{}
	err := json.Unmarshal(body, &developersResponseRoot)
	if err != nil {
		return DevelopersResponseData{}, err
	}

	if developersResponseRoot.Data.Count < 1 {
		return DevelopersResponseData{}, errors.ErrDevelopersNotFoundAPI{}
	}

	return developersResponseRoot.Data, nil
}
