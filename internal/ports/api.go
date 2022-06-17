package ports

import "net/http"

type GatewayInPortsApp interface {
	GatewayProcessRequest(response *http.Request) []byte
	GetServiceResource(serviceID string) (string, string)
}

type GatewayOutPortsApp interface {
	GetServiceResource(serviceID string) (string, string)
}
