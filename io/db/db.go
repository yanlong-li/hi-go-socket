package db

import (
	"database/sql"
	"log"
)

var db *sql.DB

//ConfigDb 配置数据库
func ConfigDb(driverName, dsn string) {
	_db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Panic(err)
	}
	db = _db
}
