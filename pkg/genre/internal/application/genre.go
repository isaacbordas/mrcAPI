package application

import (
	"mrcAPI/pkg/genre"
	"mrcAPI/pkg/genre/internal/infrastructure"
)

type Genre struct {
	r infrastructure.GenreMysqlRepository
}

func NewGenreService(r infrastructure.GenreMysqlRepository) Genre {
	return Genre{r: r}
}

func (g Genre) GetGenres() ([]genre.Genre, error) {
	return g.r.GetAllGenres()
}
