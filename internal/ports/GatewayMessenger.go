package ports

/*
	This defines all functions needed to be implemented by a service for it to be a GatewayMessenger.
	GatewayMessenger is a HTTP Client that receives a Request from the GatewayService and executes it.

	It also needs access to DB to authenticate or check authorization that a Gateway has on it.
*/
type GatewayMessengerPorts interface {
	/*
		This method receives a ServiceRequsition Form, parse it and makes another http request on private
		secured network on behalf of the user requesting a resource.

		HTTPCLIENT
	*/
	ExecuteClientRequest()

	/*
		Checks Authorization of a GatewayService against some set rules
	*/
	CheckAuthorization()
}
