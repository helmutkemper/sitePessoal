package main

import (
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func (e *MongoDBMenu) Install() (err error) {
	var installed = false

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)

	installed, err = e.verifyInstallMenu()
	if err != nil {
		util.TraceToLog()
		return
	}

	if installed == false {
		err = e.CreateTableMenu()
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.populateInitialMenu()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}
