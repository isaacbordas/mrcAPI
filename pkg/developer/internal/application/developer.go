package application

import (
	"mrcAPI/pkg/developer"
	"mrcAPI/pkg/developer/internal/infrastructure"
)

type Developer struct {
	r infrastructure.DeveloperMysqlRepository
}

func NewDeveloperService(r infrastructure.DeveloperMysqlRepository) Developer {
	return Developer{r: r}
}

func (d Developer) GetDevelopers() ([]developer.Developer, error) {
	return d.r.GetAllDevelopers()
}
