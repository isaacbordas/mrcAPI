package game

type IGame interface {
	GetGames() ([]Game, error)
	GetGame(slug string) (Game, error)
}

type Game struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
