package core

import (
	"fmt"
	"gateway/internal/adapters/models/headers"
	"gateway/internal/ports"
	"io/ioutil"
	"net/http"
	"strings"
)

type InAdapter struct {
	db ports.DBPort
}

func NewInAdapter(coms ports.DBPort) *InAdapter {
	return &InAdapter{db: coms}
}

func (ia InAdapter) GatewayProcessRequest(r *http.Request) ([]byte, headers.GatewayControl) {
	mtd := strings.ToUpper(r.Method)
	rsHeaders := headers.GatewayControl{
		ContentControl: map[string]string{
			"Content-Type": "text/plain",
		},

		CookiesControl: map[string]string{
			"Set-Cookie": "MyCookie",
		},
	}
	if mtd != "POST" {
		return []byte{}, rsHeaders
	}
	serviceReq := r.Header.Get("serviceID")
	serviceID, _ := ia.db.GetGatewayService(serviceReq)
	fmt.Println(serviceID)
	return ia.SeekService("http://localhost:5000"), rsHeaders
}

func (ia InAdapter) GetServiceResource(serviceID string) (string, string) {
	return ia.db.GetGatewayService("s")
}

func (ia InAdapter) SeekService(serviceURL string) []byte {
	client := &http.Client{}
	payload := strings.NewReader(`
		{
			"username":"myusername",
			"password":"mypassword"
		}
	`)
	req, err := http.NewRequest("GET", serviceURL, payload)
	if err != nil {
		fmt.Println("Stop processing request return error,set timeouts")
	}

	req.Header.Add("SchoolID", "test123")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Private network error, resource not found")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("There was an error parsing the body")

	}
	fmt.Println(string(body))
	return body
}
