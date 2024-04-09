package service

type Service struct {
	Shortcut
	Weather
	Schedule
	Search
}

func New() *Service {
	service := &Service{}
	return service
}
