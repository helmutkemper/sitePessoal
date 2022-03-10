package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
	"github.com/helmutkemper/util"
)

// Get (Português): Retorna o menu escolhido dentro do formato do datasource
func (e *SQLiteMenu) Get() (menu []dataformat.Menu, length int, err error) {
	var ref = make([]menuRef, 0)

	menu, err = e.getBySecondaryId("")
	if err != nil {
		util.TraceToLog()
		return
	}

	for k := range menu {
		ref = append(ref, menuRef{id: menu[k].Id, ref: &menu[k].Menu})
	}

	for {
		err = e.getReference(&ref)
		if err != nil {
			util.TraceToLog()
			return
		}

		e.clearReference(&ref)

		if len(ref) == 0 {
			break
		}
	}

	length = len(menu)

	return
}
