package httpService

import (
	"fmt"
	ports "gateway/internal/ports"
	"io"
	"net/http"
)

/*
	This Interface implements none of the ports, it uses or plugs into some ports
	It receives API
*/
type HTTPService struct {
	api ports.HTTPAPIPorts
}

func STARTHTTPSERVICE(api ports.HTTPAPIPorts) *HTTPService {
	return &HTTPService{api: api}
}

func (ht HTTPService) StartListening(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Gateway Service is listening...OK")
		ht.api.GatewayProcessRequest()
	})

	fmt.Println("GATEWAY HTTP SERVICE IS STARTING...OK")
	http.ListenAndServe(port, nil)
}
