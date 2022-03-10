package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/util"
)

func (e *MongoDBMenu) Set(id, idSecondary string, typeContent int, classroom int, text string, admin int, icon, url string, order int) (err error) {

	_, err = e.ClientMenu.InsertOne(
		e.Ctx,
		dataformat.Menu{
			Id:          id,
			IdSecondary: idSecondary,
			Text:        text,
			Admin:       admin,
			Icon:        icon,
			Url:         url,
			ItemOrder:   order,
			TypeContent: typeContent,
			Classroom:   classroom,
			Menu:        nil,
		},
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
