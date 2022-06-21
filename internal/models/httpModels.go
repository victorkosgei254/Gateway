package models

/*
	HttpHeaders are the required headers in a request for it to be considered a valid service requisition request
	ContentType should be always application/json
	ServiceID should be the predefined service ID
	Methods user does not provide this, comes by default.
	Authorization This is the API key used.
*/
type ServiceRequest struct {
	ContentType   string `json:"content-Type"`
	ServiceID     string `json:"serviceID"`
	Method        string `json:"method"`
	Authorization string `json:"authorization"`
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
