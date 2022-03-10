package main

import (
	"database/sql"
)

var Menu SQLiteMenu

type SQLiteMenu struct {
	Database *sql.DB
}
