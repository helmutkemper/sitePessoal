package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func (e *SQLiteLanguage) Connect(connectionString string, args ...interface{}) (err error) {
	e.Database, err = sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Printf("sqlite3.Connect().error: %v", err.Error())
	}
	return
}
