package main

import (
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func (e *MongoDBLanguage) Install() (err error) {
	var installed = false

	e.ClientLanguage = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionLanguage)

	installed, err = e.verifyInstallLanguage()
	if err != nil {
		util.TraceToLog()
		return
	}

	if installed == false {
		err = e.CreateTableLanguage()
		if err != nil {
			util.TraceToLog()
			return
		}

		err = e.populateInitialLanguages()
		if err != nil {
			util.TraceToLog()
			return
		}
	}

	return
}
