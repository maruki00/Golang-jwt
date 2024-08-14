package internal

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDB() *sql.DB {

	if db != nil {
		return db
	}
	db, err := sql.Open("mysql", "user:user@/go_jwt")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db

}
