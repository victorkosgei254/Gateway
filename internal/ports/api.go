package ports

import (
	"gateway/internal/adapters/models/headers"
	"net/http"
)

type GatewayInPortsApp interface {
	GatewayProcessRequest(response *http.Request) ([]byte, headers.GatewayControl)
	GetServiceResource(serviceID string) (string, string)
}

type GatewayOutPortsApp interface {
	GetServiceResource(serviceID string) (string, string)
}
