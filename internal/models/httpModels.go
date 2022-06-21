package models

import "net/http"

/*
	HttpHeaders are the required headers in a request for it to be considered a valid service requisition request
	ContentType should be always application/json
	ServiceID should be the predefined service ID
	Methods user does not provide this, comes by default.
	Authorization This is the API key used.
*/
type ServiceRequest struct {
	ContentType   string         `json:"content-Type"`
	ServiceID     string         `json:"serviceID"`
	APIKey        string         `json:"apiKey"`
	Method        string         `json:"method"`
	Authorization string         `json:"authorization"`
	Cookies       []*http.Cookie `json:"cookie"`
}

/*
	This is the Request body
*/
type ServicePayload struct {
	Payload []byte `json:"payload"`
}

/*
	Botht the Service Request and Service Payload makes a ServiceRequisistionForm
*/
type ServiceRequisitionForm struct {
	ServiceHeaders ServiceRequest `json:"serviceHeaders"`
	ServiceBody    ServicePayload `json:"serviceBody"`
}

/*
	This is a response received from the gateways service
	It is composed of
	1.The Body information for http these are the headers to be set.
	2.The Cookies to be set
	3.The Authorization to be set
	4.The Response body
*/
type ServiceResponse struct {
	Cookies       []*http.Cookie
	Authorization string
}
