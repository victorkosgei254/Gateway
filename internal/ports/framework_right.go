package ports

type DBPort interface {
	GetGatewayService(serviceID string) (string, string)
}
