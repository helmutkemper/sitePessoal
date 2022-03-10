package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBMenu) getBySecondaryId(secondaryId string) (menu []dataformat.Menu, err error) {
	var cursor *mongo.Cursor
	var menuToQuery = dataformat.Menu{IdSecondary: secondaryId}

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)
	cursor, err = e.ClientMenu.Find(e.Ctx, menuToQuery.GetIdAndSecondaryIdAsBSonQuery())
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &menu)
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}
