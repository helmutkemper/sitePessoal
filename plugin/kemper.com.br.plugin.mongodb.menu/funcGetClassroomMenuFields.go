package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBMenu) GetClassroomMenuFields() (menu []dataformat.Menu, length int, err error) {
	var cursor *mongo.Cursor

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionMenu)
	cursor, err = e.ClientMenu.Find(e.Ctx, bson.M{"classroom": 1})
	if err != nil {
		util.TraceToLog()
		return
	}

	//fixme
	err = cursor.All(e.Ctx, &menu)
	if err != nil {
		util.TraceToLog()
		return
	}

	length = len(menu)
	return
}
