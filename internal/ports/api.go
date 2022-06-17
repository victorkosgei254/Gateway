package ports

type GatewayInPortsApp interface {
	GatewayProcessRequest(requestMethod string) bool
	GetServiceResource(serviceID string) (string, string)
}

type GatewayOutPortsApp interface {
	GetServiceResource(serviceID string) (string, string)
}
