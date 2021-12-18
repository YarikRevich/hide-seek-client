package dialer

import (
	"time"

	"github.com/YarikRevich/HideSeek-Client/internal/core/statemachine"

	"github.com/YarikRevich/game-networking/pkg/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/encoding/gzip"
)

type Dialer struct {
	// locked          bool
	conn            *grpc.ClientConn
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

	conn, err := grpc.Dial("127.0.0.1:8090", opts...)
	if err != nil {
		logrus.Fatal(err)
	}

	grpc.UseCompressor(gzip.Name)

	d.conn = conn
}

func (d *Dialer) Conn() *grpc.ClientConn {
	return d.conn
}

// func (d *Dialer) Lock() {
// 	d.locked = true
// }

// func (d *Dialer) Unlock() {
// 	d.locked = false
// }

// func (d *Dialer) WaitUntilDone() {
// 	for d.locked {
// 	}
// }

func (d *Dialer) IsConnected() bool {
	if d.conn != nil {
		s := d.conn.GetState()
		return !(s == connectivity.TransientFailure || s == connectivity.Idle)
	}
	return true
}

func (d *Dialer) Reconnect() {
	select {
	case <-d.reconnectTicker.C:
		d.conn.Connect()
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
