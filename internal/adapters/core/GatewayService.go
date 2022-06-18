package core

import (
	"fmt"
	ports "gateway/internal/ports"
)

/*
	This is the implementation of a GatewaySerive. All it does can be simplified as
	1.Receives requests from client
	2.Process the requests by checking authentication or authorization
	3.From service ID obtain the location of the requested resource.
	4.Forward the Request to GatewayMessenger for execution

	It requires DataBase, GatewayMessenger
*/
type GatewayService struct {
	db  ports.DBPorts
	msg ports.GatewayMessengerPorts
}

func STARTGATEWAYSERVICE(db ports.DBPorts, msg ports.GatewayMessengerPorts) *GatewayService {
	fmt.Println("STARTING GATEWAY SERVICE...OK")
	return &GatewayService{db: db, msg: msg}

}

func (gw GatewayService) GatewayProcessRequest() {
	fmt.Println("GATEWAY SERIVE -> RECEIVED REQUEST FOR PROCESSING...OK")
	fmt.Println("GATEWAY SERIVE -> MAKING DB REQUESTS")
	fmt.Println("GATEWAY SERIVE -> Get gateway settings...OK")
	gw.db.GetGatewaySettings()
	fmt.Println("GATEWAY SERIVE -> Get gateway service config...OK")
	gw.db.GetServiceConfig()
	fmt.Println("GATEWAY SERIVE -> Get  verify API KEY...OK")
	gw.db.VerifyAPIKEY("123")
	fmt.Println("GATEWAY SERIVE -> Get user auth...OK")
	gw.db.GetUserAuth()
	fmt.Println("GATEWAY SERIVE -> EVERYTHING CHECKS OUT SENDIN REQUEST TO MESSENGER")
	gw.msg.CheckAuthorization()
	gw.msg.ExecuteClientRequest()
	fmt.Println("GATEWAY SERIVE -> CLIENT GETS RESPONSE")
}
