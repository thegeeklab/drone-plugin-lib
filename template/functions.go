package template

import (
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func loadFuncMap() template.FuncMap {
	sprigFuncs := sprig.GenericFuncMap()
	customFuncs := template.FuncMap{}

	for name, f := range customFuncs {
		if _, ok := sprigFuncs[name]; ok {
			continue
		}

		sprigFuncs[name] = f
	}

	return sprigFuncs
}
