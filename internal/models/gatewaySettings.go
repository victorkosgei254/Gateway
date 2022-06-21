package models

/*
	This model is used to configure gateway based on service ID
	First Methods for post and get are checked.
	Custom methods override some settings, which will make the Gateway perform
	a query on the authorizations
	Authorizations and authentication if false, override check api keys
*/
type GatewaySettings struct {
	ServiceID      string `json:"serviceID"`
	Post           bool   `json:"post"`
	Get            bool   `json:"get"`
	Authorization  bool   `json:"authorization"`
	Authentication bool   `json:"authentication"`
	CustomMethods  bool   `json:"customMethods"`
}
