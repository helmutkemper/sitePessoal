package main

import (
	"database/sql"
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/util"
)

// getBySecondaryId (Português): ajuda a montar o menu, navegando do nível mais externo para o mais interno.
func (e *SQLiteMenu) getBySecondaryId(secondaryId string) (menu []dataformat.Menu, err error) {
	var rows *sql.Rows

	var id string
	var idSecondary string
	var text string
	var admin int
	var icon string
	var url string
	var itemOrder int

	menu = make([]dataformat.Menu, 0)

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
				secondaryId = ?
		ORDER BY
		         itemOrder`,
		secondaryId,
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

	return
}
