package infrastructure

import (
	"encoding/json"
	"github.com/google/uuid"
	"mrcAPI/pkg/errors"
	"mrcAPI/pkg/publisher"
	"mrcAPI/pkg/system/http"
)

const publisherEndpoint = "v1/Publishers"

type PublisherImporterRepository interface {
	GetPublishers() ([]publisher.Publisher, error)
}

type PublisherImporterAPIRepository struct {
	client http.Client
}

type PublishersResponseRoot struct {
	Code   int                    `json:"code"`
	Status string                 `json:"status"`
	Data   PublishersResponseData `json:"data"`
}

type PublishersResponseData struct {
	Count      int32                 `json:"count"`
	Publishers map[string]Publishers `json:"publishers"`
}

type Publishers struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func NewPublisherImporterAPIRepository(client http.Client) PublisherImporterAPIRepository {
	return PublisherImporterAPIRepository{client: client}
}

func (i PublisherImporterAPIRepository) GetPublishers() ([]publisher.Publisher, error) {
	var publishers []publisher.Publisher

	response, err := i.client.GetRequest(publisherEndpoint, nil)
	if err != nil {
		return nil, err
	}

	transformed, errTransform := i.transformResponse(response)
	if errTransform != nil {
		return nil, errTransform
	}

	for _, pub := range transformed.Publishers {
		publishers = append(publishers, publisher.Publisher{
			UUID:  uuid.NewString(),
			ApiID: pub.ID,
			Name:  pub.Name,
		})
	}

	return publishers, err
}

func (i PublisherImporterAPIRepository) transformResponse(body []byte) (PublishersResponseData, error) {
	publishersResponseRoot := PublishersResponseRoot{}
	err := json.Unmarshal(body, &publishersResponseRoot)
	if err != nil {
		return PublishersResponseData{}, err
	}

	if publishersResponseRoot.Data.Count < 1 {
		return PublishersResponseData{}, errors.ErrPublishersNotFoundAPI{}
	}

	return publishersResponseRoot.Data, nil
}
