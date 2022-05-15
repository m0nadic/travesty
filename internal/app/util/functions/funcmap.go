package functions

import (
	"text/template"
)

var funcMap template.FuncMap = template.FuncMap{
	"faker": func() *fakerNS { return &fakerNS{} },
}

func GlobalFunctions() template.FuncMap {
	return funcMap
}
