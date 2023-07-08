package platform

type IPlatform interface {
	GetPlatforms() ([]Platform, error)
	GetPlatform(slug string) (Platform, error)
}

type Platform struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
