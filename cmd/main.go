package main

import (
	"gateway/internal/adapters/api"
	"gateway/internal/adapters/core"
	httpservice "gateway/internal/adapters/framework/left/http"
	gatewaymessenger "gateway/internal/adapters/framework/right/GatewayMessenger"
	"gateway/internal/adapters/framework/right/db"
	"gateway/internal/ports"
)

func main() {

	var db ports.GWDBPorts = db.USESQLDB("root:kosgei4182.@tcp(localhost:3306)/test_123?parseTime=true")
	var gwMessenger ports.GatewayMessengerPorts = gatewaymessenger.INITIALIZEGWMESSENGER(db)
	var gwService ports.GatewayPorts = core.STARTGATEWAYSERVICE(db, gwMessenger)
	var api ports.HTTPAPIPorts = api.INITHTTPAPI(gwService)
	gwServer := httpservice.STARTHTTPSERVICE(api)
	gwServer.StartListening(":9000")

}
