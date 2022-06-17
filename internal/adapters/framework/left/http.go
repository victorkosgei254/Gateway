package left

import (
	"fmt"
	"gateway/internal/adapters/models/headers"
	"gateway/internal/ports"
	"io"
	"net/http"
)

type HTTPAdapter struct {
	app ports.GatewayInPortsApp
}

func NewHttpAdapter(app ports.GatewayInPortsApp) *HTTPAdapter {
	return &HTTPAdapter{app: app}
}

type Service struct {
	ServiceID string            `json:"serviceID"`
	Payload   map[string]string `json:"payload"`
}

func (adpt HTTPAdapter) StartListening(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serviceResponse, serviceHeaders := adpt.app.GatewayProcessRequest(r)
		fmt.Println(serviceHeaders)
		adpt.GatewayResourceControl(w, serviceHeaders)
		io.WriteString(w, string(serviceResponse))
	})

	http.ListenAndServe(port, nil)
}

func (adpt HTTPAdapter) RedirectToRequestedResource() {
	fmt.Println("Redirecting user")
}

func (adpt HTTPAdapter) GatewayResourceControl(w http.ResponseWriter, target headers.GatewayControl) {
	//Receives headers and write them to response
	for k, v := range target.ContentControl {
		w.Header().Add(k, v)
	}

}
