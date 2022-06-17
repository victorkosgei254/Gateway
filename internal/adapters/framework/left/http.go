package left

import (
	"fmt"
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
		serviceResponse := adpt.app.GatewayProcessRequest(r)
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(serviceResponse))
	})

	http.ListenAndServe(port, nil)
}

func (adpt HTTPAdapter) RedirectToRequestedResource() {
	fmt.Println("Redirecting user")
}
