package game

type IGame interface {
	GetGames() ([]Game, error)
}

type Game struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
