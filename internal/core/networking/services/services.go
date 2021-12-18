package services

var instance *Services

type Services struct {
	base         *Base
	availability *Availability
}

func (s *Services) Availability() *Availability {
	return s.availability
}

func (s *Services) Base() *Base {
	return s.base
}

func UseServices() *Services {
	if instance == nil {
		instance = &Services{
			base:         NewBase(),
			availability: NewAvailability(),
		}
	}
	return instance
}
