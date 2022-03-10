package main

import (
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func (e *SQLiteLanguage) New() (referenceInicialized interface{}, err error) {
	//referenceInicialized = &SQLiteLanguage{}
	err = e.Connect(constants.KSQLiteConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = e.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return e, nil
}
