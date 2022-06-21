package core

import (
	"fmt"
	"gateway/internal/models"
	ports "gateway/internal/ports"
	"strings"
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
	db  ports.GWDBPorts
	msg ports.GatewayMessengerPorts
}

func STARTGATEWAYSERVICE(db ports.GWDBPorts, msg ports.GatewayMessengerPorts) *GatewayService {
	fmt.Println("STARTING GATEWAY SERVICE...OK")
	return &GatewayService{db: db, msg: msg}

}

func (gw GatewayService) GatewayProcessRequest(reqForm models.ServiceRequisitionForm) ([]byte, interface{}) {

	hed := map[string]string{
		"Content-Type": "application/json",
	}
	settings := gw.db.GetGatewaySettings(reqForm.ServiceHeaders.ServiceID)
	keyAutho := gw.db.VerifyAPIKEY(reqForm.ServiceHeaders.APIKey)

	settingsCheckSum := settings.CheckSum == "123"
	keyAuthoCheckSum := keyAutho.CheckSum == "123"
	methodCheck := strings.Contains(settings.Methods, reqForm.ServiceHeaders.Method)
	mserviceCheck := strings.Contains(keyAutho.Access, reqForm.ServiceHeaders.ServiceID)
	var resp []byte
	if methodCheck && mserviceCheck && settingsCheckSum && keyAuthoCheckSum {
		resp = gw.msg.ExecuteClientRequest(settings)
	} else {
		resp = []byte(models.ERR_MSG)
	}

	return resp, hed
}
