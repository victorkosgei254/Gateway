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

		if !adpt.app.GatewayProcessRequest(r.Method) {
			io.WriteString(w, "Bad request...OK")
			return
		}
		resiurceURL, resourcePort := adpt.app.GetServiceResource("serviceID")
		fmt.Println(resiurceURL, resourcePort)
		adpt.RedirectToRequestedResource()
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{results:"Resource obtained successfully"}`)

	})

	http.ListenAndServe(port, nil)
}

func (adpt HTTPAdapter) RedirectToRequestedResource() {
	fmt.Println("Redirecting user")
}
