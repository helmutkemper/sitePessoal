package main

import (
	dataformat "github.com/helmutkemper/kemper.com.br.module.dataformat"
)

type menuRef struct {
	id   string
	ref  *[]dataformat.Menu
	pass bool
}
