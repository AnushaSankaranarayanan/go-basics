package database

import (
	"database/sql"
	"log"
	"time"
)

// DbConn DbConn
var DbConn *sql.DB

// SetupDatabase SetupDatabase
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	DbConn.SetMaxIdleConns(4)
	DbConn.SetMaxOpenConns(4)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
