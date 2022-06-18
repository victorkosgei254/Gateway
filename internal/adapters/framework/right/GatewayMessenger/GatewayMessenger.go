package gatewaymessenger

import (
	"fmt"
	ports "gateway/internal/ports"
)

/*
This is a GatewayMessenger implementation. There can be different implementations this is just one of them.
It requires a DB so its going to be injected here.
*/
type GatewayMessenger struct {
	db ports.DBPorts
}

/*
	Creates an instance of GatewayMessenger and Injects database to it.
*/
func INITIALIZEGWMESSENGER(db ports.DBPorts) *GatewayMessenger {
	fmt.Println("MESSENGER INITIALIZED ...OK")
	return &GatewayMessenger{db: db}
}

func (msg GatewayMessenger) ExecuteClientRequest() {
	fmt.Println("GATEWAY MESSENGER - > EXECCUTING CLIENT REQUEST...OK")
}

func (msg GatewayMessenger) CheckAuthorization() {
	fmt.Println("GATEWAY MESSENGER - > CHECKING AUTHORIZATION...OK")
}
