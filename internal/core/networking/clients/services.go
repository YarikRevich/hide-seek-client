package clients

import (
	"github.com/YarikRevich/hide-seek-client/internal/core/networking/api/services_external"
	"google.golang.org/grpc"
)

type Services struct {
	client services_external.ExternalServicesServiceClient
}

func (a *Services) Connect(c *grpc.ClientConn) {
	a.client = services_external.NewExternalServicesServiceClient(c)
}

func (a *Services) GetClient() services_external.ExternalServicesServiceClient {
	return a.client
}

func NewServices() *Services {
	return new(Services)
}
