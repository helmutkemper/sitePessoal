package main

import (
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
)

func (e *MongoDBMenu) CreateTableMenu() (err error) {

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)
	return
}
