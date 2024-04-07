//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"mrcAPI/pkg/publisher"
	"mrcAPI/pkg/publisher/internal/application"
	"mrcAPI/pkg/publisher/internal/infrastructure"
	"mrcAPI/pkg/system"
)

func providePublisherMysqlRepository() (infrastructure.PublisherMysqlRepository, error) {
	panic(wire.Build(
		system.ProvideMySQL,
		infrastructure.NewPublisherMysqlRepository,
	))
}

func ProvidePublisher() (publisher.IPublisher, error) {
	panic(wire.Build(
		providePublisherMysqlRepository,
		application.NewPublisherService,
		wire.Bind(new(publisher.IPublisher), new(application.Publisher)),
	))
}
