package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func (e *MongoDBMenu) verifyInstallMenu() (installed bool, err error) {
	var cursor *mongo.Cursor
	var users []dataformat.User
	var userToQuery = dataformat.User{Id: constants.KMainUserID, Mail: constants.KMainUserMail}

	e.ClientMenu = e.Client.Database(constants.KMongoDBDatabase).Collection(constants.KMongoDBCollectionUser)

	cursor, err = e.ClientMenu.Find(e.Ctx, userToQuery.GetIdAndMailAsBSonQuery())
	if err != nil {
		util.TraceToLog()
		return
	}

	err = cursor.All(e.Ctx, &users)
	if err != nil {
		util.TraceToLog()
		return
	}

	installed = len(users) > 0
	return
}
