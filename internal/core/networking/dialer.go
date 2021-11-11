package networking

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api"
	// "github.com/YarikRevich/HideSeek-Client/internal/core/notifications"
	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"

	// "github.com/YarikRevich/game-networking/pkg/client"
	"github.com/YarikRevich/game-networking/pkg/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	_ "google.golang.org/grpc/connectivity"
)

type Dialer struct {
	conn            api.HideSeekClient
	service         *grpc.ClientConn
	reconnectTicker *time.Ticker
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

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBackoffConfig(grpc.BackoffConfig{MaxDelay: time.Second}),
	}

	service, err := grpc.Dial("127.0.0.1:8090", opts...)
	if err != nil {
		logrus.Fatal(err)
	}

	// go time.AfterFunc(time.Second*1, func() {
	// 	s := conn.GetState()
	// 	if s == connectivity.TransientFailure ||
	// 		s == connectivity.Idle || s == connectivity.Connecting {
	// 			notifications.PopUp.WriteError("Servers are offline!")
	// 			statemachine.UseStateMachine().Networking().SetState(statemachine.NETWORKING_OFFLINE)
	// 	}
	// })

	d.service = service
	d.conn = api.NewHideSeekClient(service)
}

func (d *Dialer) Conn() api.HideSeekClient {
	return d.conn
}

func (d *Dialer) IsConnected() bool {
	if d.conn != nil {
		s := d.service.GetState()
		return !(s == connectivity.TransientFailure || s == connectivity.Idle)
	}
	return true
}

func (d *Dialer) Reconnect() {
	select {
	case <-d.reconnectTicker.C:
		d.service.Connect()
	default:
	}
}

func NewDialer() *Dialer {
	return &Dialer{
		reconnectTicker: time.NewTicker(time.Millisecond * 500),
	}
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
