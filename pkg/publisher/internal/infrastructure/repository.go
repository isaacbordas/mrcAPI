package infrastructure

import (
	"database/sql"
	"mrcAPI/pkg/publisher"
)

type PublisherMysqlRepository struct {
	db *sql.DB
}

func NewPublisherMysqlRepository(db *sql.DB) PublisherMysqlRepository {
	return PublisherMysqlRepository{db: db}
}

func (g PublisherMysqlRepository) GetAllPublishers() ([]publisher.Publisher, error) {
	var publishers []publisher.Publisher

	query := "SELECT p.publisher_uuid, p.publisher_id, p.publisher_name FROM publishers p ORDER BY p.publisher_name"

	result, err := g.db.Query(query)
	if err != nil {
		return publishers, err
	}

	var publisherUuid, publisherName string
	var publisherID int32
	for result.Next() {
		err = result.Scan(&publisherUuid, &publisherID, &publisherName)
		if err != nil {
			result.Close()
			return publishers, err
		}
		publishers = append(publishers, publisher.Publisher{
			UUID:  publisherUuid,
			ApiID: publisherID,
			Name:  publisherName,
		})
	}

	return publishers, nil
}
