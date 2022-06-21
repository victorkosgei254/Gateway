package httpService

import (
	"fmt"
	"gateway/internal/models"
	ports "gateway/internal/ports"
	"io"
	"io/ioutil"
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
		fmt.Println()
		serviceHeaders := models.ServiceRequest{
			ContentType:   r.Header.Get("Content-Type"),
			ServiceID:     r.Header.Get("ServiceID"),
			Method:        r.Method,
			Authorization: r.Header.Get("Authorization"),
		}

		body, _ := ioutil.ReadAll(r.Body)

		servicePayload := models.ServicePayload{
			Payload: body,
		}
		reqForm := models.ServiceRequisitionForm{
			ServiceHeaders: serviceHeaders,
			ServiceBody:    servicePayload,
		}
		// io.WriteString(w, "Gateway Service is listening...OK")
		res, _ := ht.api.GatewayProcessRequest(reqForm)
		w.Header().Add("Content-Type", "application/json")

		fmt.Println("*************")
		fmt.Println(string(res))
		io.WriteString(w, string(res))

	})

	fmt.Println("GATEWAY HTTP SERVICE IS STARTING...OK")
	http.ListenAndServe(port, nil)
}
