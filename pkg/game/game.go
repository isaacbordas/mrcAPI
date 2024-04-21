package game

type IGame interface {
	GetGameByName(gameName string) ([]Game, error)
}

type Game struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Platform int32  `json:"platform"`
	Region   int32  `json:"region"`
}
