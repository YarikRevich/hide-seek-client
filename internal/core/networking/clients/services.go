package clients

import (
	"github.com/YarikRevich/HideSeek-Client/internal/core/networking/api/services_external"
	"google.golang.org/grpc"
)

type Services struct {
	client services_external.ExternalServiceClient
}

func (a *Services) Connect(c *grpc.ClientConn) {
	a.client = services_external.NewExternalServiceClient(c)
}

func (a *Services) GetClient() services_external.ExternalServiceClient {
	return a.client
}

func NewServices() *Services {
	return new(Services)
}
