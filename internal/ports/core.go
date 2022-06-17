package ports

type GatewayInPorts interface {
	GatewayProcessRequest(requestMethod string) bool
	GetServiceResource(serviceID string) (string, string)
}

type GatewayOutPorts interface {
	GetServiceResource(serviceID string) (string, string)
}
