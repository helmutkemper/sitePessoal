package main

import (
	"github.com/helmutkemper/util"
)

func (e *MongoDBLanguage) Close() (err error) {
	err = e.Client.Disconnect(e.Ctx)
	if err != nil {
		util.TraceToLog()
	}
	return
}
