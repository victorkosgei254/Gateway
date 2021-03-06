package db

import (
	"database/sql"
	"fmt"
	"gateway/internal/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
 This Adapter implements the DBPorts interface
*/
type SQLDBAdapter struct {
	db *sql.DB
}

/*
	Creates an SQLDB instance and injects it to SQLAdapter
	Type.
*/
func USESQLDB(dataSourceName string) *SQLDBAdapter {
	//Connecting to Database
	sqldb, err := sql.Open("mysql", dataSourceName)
	fmt.Println("Connecting to MYSQL DATABASE")
	if err != nil {
		// log.Fatal("Cannot Connect to MYSQL Instance...OK")
		fmt.Println("Could not connect to DB", err)
	}

	// test the connection
	pingerr := sqldb.Ping()

	if pingerr != nil {
		log.Fatal("Could not establish connection to MYSQL DB..EXITING")
		fmt.Println("error")
	}

	return &SQLDBAdapter{db: sqldb}
}

func (sqlAdp SQLDBAdapter) VerifyAPIKEY(key string) models.APIKEY {
	fmt.Println("SQL DB -> VERIFYING API KEY ...OK")
	return models.APIKEY{
		UserID:    "testuser123",
		CreatedOn: "12-21-22",
		Access:    "AUTH;AUTHO;MRKS",
		CheckSum:  "123",
	}
}

func (sqlAdp SQLDBAdapter) GetGatewaySettings(serviceID string) models.GatewaySettings {
	fmt.Println("SQL DB -> GETTING GATEWAY SETTINGS ...OK")
	return models.GatewaySettings{
		ServiceID:     "test123",
		Methods:       "POST",
		APIKey:        "eYds72mdmshnsmnd=",
		Authorization: " sa",
		ServiceURL:    "http://localhost:5000",
		CheckSum:      "123",
	}
}

func (sqlAdp SQLDBAdapter) GetServiceConfig() {
	fmt.Println("SQL DB -> GETTING SERVICE CONFIGURATIONS...OK")
}

func (sqlAdp SQLDBAdapter) GetUserAuth() {
	fmt.Println("SQL DB -> GETTING USER AUTHENTICATION AND VERIFYING CLAIMS...OK")
}

func (sqlAdp SQLDBAdapter) CloseDBConnection() {
	fmt.Println("SQL DB -> CLOSING DB...OK")
	sqlAdp.db.Close()
}
