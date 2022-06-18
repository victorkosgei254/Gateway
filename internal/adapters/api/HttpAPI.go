package api

import (
	"fmt"
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

func (api HTTPApi) GatewayProcessRequest() {
	fmt.Println("HTTP API SERVICE -> FROM HTTP API....MAKING A CALL TO CORE, GATEWAYPROCESS REQUEST...OK")
	api.gw.GatewayProcessRequest()
}
