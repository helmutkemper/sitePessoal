package main

import (
	"github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/util"
)

func (e *MongoDBUser) MailExists(mail string) (found bool, err error) {
	var user dataformat.User
	user, err = e.GetByEmail(mail)
	if err != nil {
		util.TraceToLog()
		return
	}

	found = user.Mail != ""
	return
}
