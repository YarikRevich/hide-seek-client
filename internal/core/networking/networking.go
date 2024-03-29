package networking

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/clients"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/dialer"
)

var instance *Networking

type Networking struct {
	dialer  *dialer.Dialer
	clients *clients.Clients
}

func (n *Networking) Dialer() *dialer.Dialer {
	return n.dialer
}

func (n *Networking) Clients() *clients.Clients {
	return n.clients
}

func (n *Networking) Init() {
	n.dialer.Dial()
	n.clients.Base().Connect(n.dialer.GetServerConn())
	n.clients.Services().Connect(n.dialer.GetServicesConn())
}

func UseNetworking() *Networking {
	if instance == nil {
		instance = &Networking{
			dialer:  dialer.NewDialer(),
			clients: clients.UseClients(),
		}
	}
	return instance
}
