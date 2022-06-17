package right

import "database/sql"

type DBAdapter struct {
	db *sql.DB
}

func NewDBAdapter(databaseName, databaseSource string) *DBAdapter {
	return &DBAdapter{}
}

func (db DBAdapter) GetGatewayService(serviceID string) (string, string) {

	return "service.localhost:port", "service/route"
}
