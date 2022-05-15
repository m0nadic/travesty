package functions

import (
	"github.com/jaswdr/faker"
	"text/template"
)

var funcMap template.FuncMap = template.FuncMap{
	"name":      func() string { return faker.New().Person().Name() },
	"firstName": func() string { return faker.New().Person().FirstName() },
	"lastName":  func() string { return faker.New().Person().LastName() },
	"uuid":      func() string { return faker.New().UUID().V4() },
}

func GlobalFunctions() template.FuncMap {
	return funcMap
}
