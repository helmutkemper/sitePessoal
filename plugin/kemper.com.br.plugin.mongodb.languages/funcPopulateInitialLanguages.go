package main

import (
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func (e *MongoDBLanguage) populateInitialLanguages() (err error) {
	err = e.Set(constants.KInstallLanguageID, constants.KInstallLanguageName)
	if err != nil {
		util.TraceToLog()
	}
	return
}
