package main

import (
	"database/sql"
	"github.com/helmutkemper/util"
)

func (e *SQLiteMenu) verifyInstallMenu() (installed bool, err error) {
	var rows *sql.Rows

	rows, err = e.Database.Query(`
		SELECT
      count(*)
		FROM
      sqlite_master
		WHERE
      type=? AND
      name=?;
    `,
		"table",
		"menu",
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	installed = count == 1

	return
}
