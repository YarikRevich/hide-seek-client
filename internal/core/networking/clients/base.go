package clients

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/server_external"
	"google.golang.org/grpc"
)

type Base struct {
	client server_external.ExternalServiceClient
}

func (b *Base) Connect(c *grpc.ClientConn) {
	server_external.NewExternalServiceClient(c)
}

func (b *Base) GetClient() server_external.ExternalServiceClient {
	return b.client
}

func NewBase() *Base {
	return new(Base)
}
