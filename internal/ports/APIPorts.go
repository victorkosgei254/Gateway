package ports

import "gateway/internal/models"

/*
	API Ports this is a middle layer that connects the Core with the drivers and also Driven.
	This ports holds all the ports that Gateway provides
	In later versions, it also include API for Loggers and Publishers
*/
type HTTPAPIPorts interface {
	GatewayProcessRequest(reqFrom models.ServiceRequisitionForm) ([]byte, interface{})
}
