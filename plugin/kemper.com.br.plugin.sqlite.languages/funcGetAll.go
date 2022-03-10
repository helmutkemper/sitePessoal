package main

import (
	"database/sql"
	"errors"
	"github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"log"
)

func (e *SQLiteLanguage) GetAll() (languagues []dataformat.Languages, length int, err error) {
	var rows *sql.Rows

	languagues = make([]dataformat.Languages, 0)

	rows, err = e.Database.Query(
		`
			SELECT
			    id,
				name
			FROM
				language
			ORDER BY 
				name
		`,
	)
	if err != nil {
		log.Printf("SQLiteLanguage.GetAll().error: %v", err.Error())
		return
	}

	var id string
	var name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Printf("SQLiteLanguage.GetAll().error: %v", err.Error())
			return
		}

		languagues = append(
			languagues,
			dataformat.Languages{
				Id:   id,
				Name: name,
			},
		)
	}

	length = len(languagues)

	if length == 0 {
		err = errors.New(constants.KErrorLanguageTableEmpty)
		log.Printf("SQLiteLanguage.GetAll().error: %v", err.Error())
	}

	return
}
