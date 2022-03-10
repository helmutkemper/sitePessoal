package main

import (
	"database/sql"
	"errors"
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/util"
)

func (e *SQLiteMenu) GetClassroomMenuFields() (menu []dataformat.Menu, length int, err error) {
	var rows *sql.Rows

	var id string
	var idSecondary string
	var text string
	var admin int
	var icon string
	var url string
	var itemOrder int

	menu = make([]dataformat.Menu, 0)

	if e.Database == nil {
		panic(errors.New("database menu is nil"))
	}

	rows, err = e.Database.Query(
		`
		SELECT
		       id,
		       secondaryId,
		       text,
		       admin,
		       icon,
		       url,
		       itemOrder
		FROM
		     menu
		WHERE
				classroom = 1
		ORDER BY
		         itemOrder`,
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	for rows.Next() {
		err = rows.Scan(&id, &idSecondary, &text, &admin, &icon, &url, &itemOrder)
		if err != nil {
			util.TraceToLog()
			return
		}

		menu = append(
			menu,
			dataformat.Menu{
				Id:          id,
				IdSecondary: idSecondary,
				Text:        text,
				Admin:       admin,
				Icon:        icon,
				Url:         url,
				ItemOrder:   itemOrder,
				Menu:        make([]dataformat.Menu, 0),
			},
		)
	}

	length = len(menu)
	return
}
