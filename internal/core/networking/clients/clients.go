package clients

var instance *Clients

type Clients struct {
	base     *Base
	services *Services
}

func (s *Clients) Services() *Services {
	return s.services
}

func (s *Clients) Base() *Base {
	return s.base
}

func UseClients() *Clients {
	if instance == nil {
		instance = &Clients{
			base:     NewBase(),
			services: NewServices(),
		}
	}
	return instance
}
