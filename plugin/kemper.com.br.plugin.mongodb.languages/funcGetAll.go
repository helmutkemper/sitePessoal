package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBLanguage) GetAll() (languagues []dataformat.Languages, length int, err error) {
	var cursor *mongo.Cursor

	e.ClientLanguage = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionLanguage)
	cursor, err = e.ClientLanguage.Find(e.Ctx, bson.M{})
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &languagues)
	if err != nil {
		util.TraceToLog()
		return
	}

	length = len(languagues)
	return
}
