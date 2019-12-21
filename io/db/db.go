package db

import (
	"database/sql"
	"log"
)

var db *sql.DB

func ConfigDb(driverName, dsn string) {
	_db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Panic(err)
	}
	db = _db
}
