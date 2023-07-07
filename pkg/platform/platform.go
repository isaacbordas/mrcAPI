package platform

type IPlatform interface {
	GetPlatforms() ([]Platform, error)
}

type Platform struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
