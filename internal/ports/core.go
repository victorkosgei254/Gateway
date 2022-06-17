package ports

import (
	"gateway/internal/adapters/models/headers"
	"net/http"
)

type GatewayInPorts interface {
	GatewayProcessRequest(requestMethod *http.Request) ([]byte, headers.GatewayControl)
	GetServiceResource(serviceID string) (string, string)
}

type GatewayOutPorts interface {
	GetServiceResource(serviceID string) (string, string)
}
