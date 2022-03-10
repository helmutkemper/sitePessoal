package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBLanguage) verifyInstallLanguage() (installed bool, err error) {
	var cursor *mongo.Cursor
	var languagues []dataformat.Languages

	//fixme: toBSon()
	cursor, err = e.ClientLanguage.Find(e.Ctx, bson.M{"_id": constants.KInstallLanguageID, "name": constants.KInstallLanguageName})
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &languagues)
	if err != nil {
		util.TraceToLog()
		return
	}

	installed = len(languagues) > 0
	return
}
