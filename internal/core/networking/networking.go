package networking

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/dialer"
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/services"
)

var instance *Networking

type Networking struct {
	dialer   *dialer.Dialer
	services *services.Services
}

func (n *Networking) Dialer() *dialer.Dialer {
	return n.dialer
}

func (n *Networking) Services() *services.Services {
	return n.services
}

func UseNetworking() *Networking {
	if instance == nil {
		instance = &Networking{
			dialer:   dialer.NewDialer(),
			services: services.UseServices(),
		}
	}
	return instance
}
