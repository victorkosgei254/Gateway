package ports

import "net/http"

type GatewayInPorts interface {
	GatewayProcessRequest(requestMethod *http.Request) []byte
	GetServiceResource(serviceID string) (string, string)
}

type GatewayOutPorts interface {
	GetServiceResource(serviceID string) (string, string)
}
