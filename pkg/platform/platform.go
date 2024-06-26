package platform

type IPlatform interface {
	GetPlatforms() ([]Platform, error)
	GetPlatform(slug string) (Platform, error)
}

type Platform struct {
	UUID  string `json:"uuid"`
	ApiID int32  `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
}
