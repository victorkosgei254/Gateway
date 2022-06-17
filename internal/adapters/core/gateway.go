package core

import (
	"gateway/internal/ports"
	"strings"
)

type InAdapter struct {
	db ports.DBPort
}

func NewInAdapter(coms ports.DBPort) *InAdapter {
	return &InAdapter{db: coms}
}

func (ia InAdapter) GatewayProcessRequest(requestMethod string) bool {
	mtd := strings.ToUpper(requestMethod)
	return mtd == "POST"
}

func (ia InAdapter) GetServiceResource(serviceID string) (string, string) {
	return ia.db.GetGatewayService("s")
}
