package main

import (
	"database/sql"
	"github.com/helmutkemper/util"
)

func (e *SQLiteMenu) Set(id, idSecondary string, typeContent int, classroom int, text string, admin int, icon, url string, order int) (err error) {
	var statement *sql.Stmt
	statement, err = e.Database.Prepare(
		`INSERT INTO main.menu (id, secondaryId, typeContent, classroom, text, admin, icon, url, itemOrder) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	_, err = statement.Exec(id, idSecondary, typeContent, classroom, text, admin, icon, url, order)
	if err != nil {
		util.TraceToLog()
		return
	}
	return
}
