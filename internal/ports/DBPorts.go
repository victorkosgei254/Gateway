package ports

/*
	This interface defines what Kind of DB ports are
	available to the  core and what kind of Action it can perform on the
	Database.

	Whether SQL, MongoDB, InMEM, PostgreSQL or any other
	the services for these Databases should implement this
	interface in order for it to interact with the core.

	For the Gateway Service, Database Operations needed are as follows
	1.Authenticate API Key
	2.Check request headers if its okay
	3.Get Gateway settings from the DB and implement them.
	4. Handles Authentication user authentication.

*/
type DBPorts interface {
	/*
		VerifyAPIKEY
		Receives APIKEY from client, for http from Headers (Authorization)
		The APIKEY has the following
		APIKEY,UUID,Authorization,HASH
		UUID is the user ID.


	*/
	VerifyAPIKEY(key string)
	/*
		Retrieve Gateway settings, i.e
		ALLOW METHODS,ALLOW SERVICES,
	*/
	GetGatewaySettings()

	/*
		Gets Service Configurations, i.e
		AUTH Service Config{
			APIKEY: Allow NULL
			METHOD : POST
			BODY: Requires Username and Password
			SERVICEURL : http://localhost:9009
		}
	*/
	GetServiceConfig()

	/*
		Performs User authentication

	*/
	GetUserAuth()
}
