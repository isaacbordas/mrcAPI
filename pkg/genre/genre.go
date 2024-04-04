package genre

type IGenre interface {
	GetGenres() ([]Genre, error)
}

type Genre struct {
	UUID  string `json:"uuid"`
	ApiID int32  `json:"id"`
	Name  string `json:"name"`
}
