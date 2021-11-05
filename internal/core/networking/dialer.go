package networking

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"
	"github.com/YarikRevich/game-networking/pkg/client"
	"github.com/YarikRevich/game-networking/pkg/config"
	"github.com/sirupsen/logrus"
)

type Dialer struct {
	conn client.Dialer
}

func (d *Dialer) Dial() {
	c := config.Config{
		Port: "8090",
	}
	switch statemachine.UseStateMachine().Dial().GetState() {
	case statemachine.DIAL_LAN:
		c.IP = "127.0.0.1"
	case statemachine.DIAL_WAN:
		c.IP = "127.0.0.1"
	}

	conn, err := client.Dial(c)
	if err != nil {
		logrus.Fatal(err)
	}
	d.conn = conn
}

func (d *Dialer) Conn() client.Dialer {
	return d.conn
}

func NewDialer() *Dialer{
	return new(Dialer)
}
// // if err != nil {
// 		// 	applyer.ApplyMiddlewares(
// 		// 		statemachine.UseStateMachine().Networking().SetState(networking.OFFLINE),
// 		// 		networkingmiddleware.UseNetworkingMiddleware,
// 		// 	)
// 		// }

// 		// instance = d

// 		// go func() {
// 		// 	sc := make(chan os.Signal, 1)
// 		// 	signal.Notify(sc, os.Interrupt)
// 		// 	for range sc {
// 		// 		if err := instance.Close(); err != nil {
// 		// 			logrus.Fatal(err)
// 		// 		}
// 		// 	}
// 		// }()
