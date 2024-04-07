package application

import (
	"mrcAPI/pkg/publisher"
	"mrcAPI/pkg/publisher/internal/infrastructure"
)

type Publisher struct {
	r infrastructure.PublisherMysqlRepository
}

func NewPublisherService(r infrastructure.PublisherMysqlRepository) Publisher {
	return Publisher{r: r}
}

func (p Publisher) GetPublishers() ([]publisher.Publisher, error) {
	return p.r.GetAllPublishers()
}
