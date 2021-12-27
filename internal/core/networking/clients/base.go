package clients

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"google.golang.org/grpc"
)

type Base struct {
	client server_external.ExternalServerServiceClient
}

func (b *Base) Connect(c *grpc.ClientConn) {
	server_external.NewExternalServerServiceClient(c)
}

func (b *Base) GetClient() server_external.ExternalServerServiceClient {
	return b.client
}

func NewBase() *Base {
	return new(Base)
}
