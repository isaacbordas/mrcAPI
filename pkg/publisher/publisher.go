package publisher

type IPublisher interface {
	GetPublishers() ([]Publisher, error)
}

type Publisher struct {
	UUID  string `json:"uuid"`
	ApiID int32  `json:"id"`
	Name  string `json:"name"`
}
