package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/util"
)

func (e *MongoDBLanguage) Set(id, name string) (err error) {

	_, err = e.ClientLanguage.InsertOne(
		e.Ctx,
		dataformat.Languages{
			Id:   id,
			Name: name,
		},
	)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
