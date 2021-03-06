package user

import (
	"errors"
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	systemDatasource "github.com/helmutkemper/kemper.com.br/businessRules/system/datasource"
	"log"
)

func (e *BusinessRules) getPasswordByEmail(mail string) (password string, err error) {
	var userFromDatasource dataformat.User
	var matched bool

	matched, err = e.verifyMailSyntax(mail)
	if err != nil {
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	if matched == false {
		err = errors.New(constants.KErrorEmailValidSintax)
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	e.DataSource = systemDatasource.Linker.GetReferenceFromUser()
	userFromDatasource, err = e.DataSource.GetByEmail(mail)
	if err != nil {
		log.Printf("user.getPasswordByEmail().error: %v", err.Error())
		return
	}

	password = userFromDatasource.Password
	return
}
