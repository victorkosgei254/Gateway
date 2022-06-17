package app

import (
	"fmt"
	"gateway/internal/ports"
	"net/http"
)

type AppAdapter struct {
	core    ports.GatewayInPorts
	comsOut ports.GatewayOutPorts
}

func NewAppAdapter(core ports.GatewayInPorts, coms ports.GatewayOutPorts) *AppAdapter {
	return &AppAdapter{core: core, comsOut: coms}
}

func (app AppAdapter) GatewayProcessRequest(requestMethod *http.Request) []byte {
	return app.core.GatewayProcessRequest(requestMethod)
}

func (app AppAdapter) GetServiceResource(serviceID string) (string, string) {
	fmt.Println("Getting Service Resource")
	return app.comsOut.GetServiceResource("s")
}

type AppOutAdapter struct {
	db ports.DBPort
}

func NewCoreToOutsideAdapter(db ports.DBPort) *AppOutAdapter {
	return &AppOutAdapter{db: db}
}
func (app AppOutAdapter) GetGatewayService(serviceID string) (string, string) {

	return "Serive.localhost.someport", "serive.url"
}

func (app AppOutAdapter) GetServiceResource(serviceID string) (string, string) {
	return app.db.GetGatewayService("s")
}
