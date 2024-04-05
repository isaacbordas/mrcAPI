package infrastructure

import (
	"database/sql"
	"mrcAPI/pkg/developer"
)

type DeveloperMysqlRepository struct {
	db *sql.DB
}

func NewDeveloperMysqlRepository(db *sql.DB) DeveloperMysqlRepository {
	return DeveloperMysqlRepository{db: db}
}

func (g DeveloperMysqlRepository) GetAllDevelopers() ([]developer.Developer, error) {
	var developers []developer.Developer

	query := "SELECT g.developer_uuid, g.developer_id, g.developer_name FROM developers g ORDER BY g.developer_name"

	result, err := g.db.Query(query)
	if err != nil {
		return developers, err
	}

	var developerUuid, developerName string
	var developerID int32
	for result.Next() {
		err = result.Scan(&developerUuid, &developerID, &developerName)
		if err != nil {
			result.Close()
			return developers, err
		}
		developers = append(developers, developer.Developer{
			UUID:  developerUuid,
			ApiID: developerID,
			Name:  developerName,
		})
	}

	return developers, nil
}
