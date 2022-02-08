package networking

import (
	"os"
	"os/exec"
	"strings"

	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/services_external"
	"github.com/YarikRevich/hide-seek-client/internal/core/statemachine"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/encoding/gzip"
)

type NetworkingManager struct {
	LANProcess    *os.Process
	WAN, LAN      *grpc.ClientConn
	ServerClient  *server_external.ExternalServerServiceClient
	ServiceClient *services_external.ExternalServicesServiceClient
}

func (nm *NetworkingManager) dial(port string) *grpc.ClientConn {
	var endpoint strings.Builder

	switch statemachine.Dial.GetState() {
	case statemachine.DIAL_LAN:
		endpoint.WriteString("127.0.0.1:")
	case statemachine.DIAL_WAN:
		endpoint.WriteString("127.0.0.1:")
	}
	endpoint.WriteString(port)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	dst, err := grpc.Dial(endpoint.String(), opts...)
	if err != nil {
		logrus.Fatal(err)
	}

	return dst
}

func (nm *NetworkingManager) StartLANServer() {
	cmd := exec.Command("")
	nm.LANProcess = cmd.Process
	//TODO: start LAN server process

	nm.LAN = nm.dial("8090")
	nm.CreateClients(nm.LAN)
}

func (nm *NetworkingManager) StopLANServer() {
	if err := nm.WAN.Close(); err != nil {
		logrus.Fatalln(err)
	}
	if err := nm.LANProcess.Kill(); err != nil {
		logrus.Fatalln(err)
	}
}

func (nm *NetworkingManager) DialWANServer() {
	nm.WAN = nm.dial("8099")
	nm.CreateClients(nm.WAN)
}

func (nm *NetworkingManager) CloseWANServer() {
	if err := nm.WAN.Close(); err != nil {
		logrus.Fatalln(err)
	}
}

func (nm *NetworkingManager) IsWANConnected() bool {
	state := nm.WAN.GetState()
	return !(state == connectivity.TransientFailure || state == connectivity.Idle)
}

func (nm *NetworkingManager) IsLANConnected() bool {
	state := nm.LAN.GetState()
	return !(state == connectivity.TransientFailure || state == connectivity.Idle)
}

func (nm *NetworkingManager) CreateClients(conn *grpc.ClientConn) {
	server_external.NewExternalServerServiceClient(conn)
	services_external.NewExternalServicesServiceClient(conn)
}

func NewNetworkingManager() *NetworkingManager {
	grpc.UseCompressor(gzip.Name)
	return new(NetworkingManager)
}
