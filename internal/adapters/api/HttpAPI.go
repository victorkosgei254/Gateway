package api

import (
	"fmt"
	"gateway/internal/models"
	ports "gateway/internal/ports"
)

/*
	This Implements all API methods
	Receives Gateway
*/
type HTTPApi struct {
	gw ports.GatewayPorts
}

func INITHTTPAPI(gw ports.GatewayPorts) *HTTPApi {
	fmt.Println("INITIALIZING API ...OK")
	return &HTTPApi{gw: gw}
}

func (api HTTPApi) GatewayProcessRequest(reqForm models.ServiceRequisitionForm) ([]byte, interface{}) {
	fmt.Println("HTTP API SERVICE -> FROM HTTP API....MAKING A CALL TO CORE, GATEWAYPROCESS REQUEST...OK")
	return api.gw.GatewayProcessRequest(reqForm)
}
