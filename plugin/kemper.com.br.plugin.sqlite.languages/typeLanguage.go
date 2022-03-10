package main

import (
	"database/sql"
)

var Language SQLiteLanguage

type SQLiteLanguage struct {
	Database *sql.DB
}
