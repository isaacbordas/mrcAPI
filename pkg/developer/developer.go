package developer

type IDeveloper interface {
	GetDevelopers() ([]Developer, error)
}

type Developer struct {
	UUID  string `json:"uuid"`
	ApiID int32  `json:"id"`
	Name  string `json:"name"`
}
