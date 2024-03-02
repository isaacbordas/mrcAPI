package application

import (
	"fmt"
	"mrcAPI/pkg/importer/internal/infrastructure"
)

type Importer struct {
	a infrastructure.ImporterAPIRepository
	r infrastructure.ImporterMysqlRepository
}

func NewImporterService(a infrastructure.ImporterAPIRepository, r infrastructure.ImporterMysqlRepository) Importer {
	return Importer{a: a, r: r}
}

func (i Importer) ImportPlatforms() error {
	platformsAPI, errAPI := i.a.GetPlatforms()
	if errAPI != nil {
		return errAPI
	}

	fmt.Println(platformsAPI)
	panic("")

	return i.r.PersistPlatforms(platformsAPI)
}
