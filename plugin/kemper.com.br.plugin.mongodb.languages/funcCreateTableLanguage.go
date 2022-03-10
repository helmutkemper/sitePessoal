package main

import (
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
)

func (e *MongoDBLanguage) CreateTableLanguage() (err error) {

	e.ClientLanguage = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionLanguage)
	return
}
