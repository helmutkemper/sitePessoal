package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
)

func (e *MongoDBMenu) getReference(ref *[]menuRef) (err error) {
	var menu []dataformat.Menu

	if ref == nil {
		return
	}

	for refKey := range *ref {
		menu, err = e.getBySecondaryId((*ref)[refKey].id)
		*(*ref)[refKey].ref = menu
		(*ref)[refKey].pass = true
		for menuKey := range menu {
			if menu[menuKey].IdSecondary != "" {
				*ref = append(*ref, menuRef{id: menu[menuKey].Id, ref: &menu[menuKey].Menu})
			}
		}
	}

	return
}
