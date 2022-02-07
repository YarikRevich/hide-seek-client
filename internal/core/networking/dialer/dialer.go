package dialer

import (
	"strings"
	"time"

	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/encoding/gzip"
)

type Dialer struct {
	server_conn, services_conn *grpc.ClientConn
	reconnectTicker            *time.Ticker
}

func (d *Dialer) dialServer() {
	var endpoint strings.Builder

	switch statemachine.Dial.GetState() {
	case statemachine.DIAL_LAN:
		endpoint.WriteString("127.0.0.1:")
	case statemachine.DIAL_WAN:
		endpoint.WriteString("127.0.0.1:")
	}
	endpoint.WriteString("8080")

	opts := []grpc.DialOption{grpc.WithInsecure()}

	server_conn, err := grpc.Dial(endpoint.String(), opts...)
	if err != nil {
		logrus.Fatal(err)
	}

	grpc.UseCompressor(gzip.Name)

	d.server_conn = server_conn
}
func (d *Dialer) dialServices() {
	var endpoint strings.Builder

	switch statemachine.Dial.GetState() {
	case statemachine.DIAL_LAN:
		endpoint.WriteString("127.0.0.1:")
	case statemachine.DIAL_WAN:
		endpoint.WriteString("127.0.0.1:")
	}
	endpoint.WriteString("8099")

	opts := []grpc.DialOption{grpc.WithInsecure()}

	services_conn, err := grpc.Dial(endpoint.String(), opts...)
	if err != nil {
		logrus.Fatal(err)
	}

	grpc.UseCompressor(gzip.Name)

	d.services_conn = services_conn
}

func (d *Dialer) Dial() {
	d.dialServer()
	d.dialServices()
}

func (d *Dialer) GetServerConn() *grpc.ClientConn {
	return d.server_conn
}
func (d *Dialer) GetServicesConn() *grpc.ClientConn {
	return d.services_conn
}

func (d *Dialer) IsServerClientConnected() bool {
	if d.server_conn != nil {
		s := d.server_conn.GetState()
		return !(s == connectivity.TransientFailure || s == connectivity.Idle)
	}
	return true
}

func (d *Dialer) IsServicesClientConnected() bool {
	if d.services_conn != nil {
		s := d.services_conn.GetState()
		return !(s == connectivity.TransientFailure || s == connectivity.Idle)
	}
	return true
}

func (d *Dialer) ReconnectServerClient() {
	select {
	case <-d.reconnectTicker.C:
		d.server_conn.Connect()
	default:
	}
}

func (d *Dialer) ReconnectServicesClient() {
	select {
	case <-d.reconnectTicker.C:
		d.services_conn.Connect()
	default:
	}
}

func NewDialer() *Dialer {
	return &Dialer{
		reconnectTicker: time.NewTicker(time.Millisecond * 500),
	}
}
