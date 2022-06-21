package models

/*
	API key models, returns a struct of an API containing the API Key authorizations
*/

type APIKEY struct {
	UserID    string `json:"userid"`
	CreatedOn string `json:"createdON"`
	Access    string `json:"access"`
	CheckSum  string `json:"checksum"`
}
