package gatewaymessenger

import (
	"fmt"
	"gateway/internal/models"
	ports "gateway/internal/ports"
	"io/ioutil"
	"net/http"
)

/*
This is a GatewayMessenger implementation. There can be different implementations this is just one of them.
It requires a DB so its going to be injected here.
*/
type GatewayMessenger struct {
	db ports.GWDBPorts
}

/*
	Creates an instance of GatewayMessenger and Injects database to it.
*/
func INITIALIZEGWMESSENGER(db ports.GWDBPorts) *GatewayMessenger {
	fmt.Println("MESSENGER INITIALIZED ...OK")
	return &GatewayMessenger{db: db}
}

func (msg GatewayMessenger) ExecuteClientRequest(resource models.GatewaySettings) []byte {
	fmt.Println("GATEWAY MESSENGER - > EXECCUTING CLIENT REQUEST...OK")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, resource.ServiceURL, nil)

	if err != nil {
		fmt.Println(err)
		return []byte(`
		"msg":"Error processing Request"
		`)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte(`
		"msg":"Error processing Request"
		`)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte(`
		"msg":"Error processing Request"
		`)
	}

	return body
}

func (msg GatewayMessenger) CheckAuthorization() {
	fmt.Println("GATEWAY MESSENGER - > CHECKING AUTHORIZATION...OK")
}
