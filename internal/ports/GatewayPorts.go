package ports

/*

This interface provide ports for driver actors to interact with the GatewayService.
Actors in this case is the http service.
Although HTTP is the only service now, In future release, another service i.e GRPC can be used instead,
therefore there is also a need for an APPLICATION layer to transform these request into models that can be used by the
Gateway
*/
type GatewayPorts interface {
	GatewayProcessRequest()
}
