package main

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func (e *MongoDBUser) New() (referenceInicialized interface{}, err error) {
	err = e.Connect(constants.KMongoDBConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return e, nil
}
