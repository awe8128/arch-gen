package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

func NewFuncTemplate(method, name string, params, out map[string]config.Property, isCtx bool, isRepo bool) string {

	fn := capitalize(method) + capitalize(name)
	p := GetParams(params, isCtx)
	r := GetReturnValues(name, out)

	if isRepo {
		return fmt.Sprintf(`%s %s %s`, fn, p, r)
	}

	return fmt.Sprintf(`func %s %s %s`, fn, p, r)
}

func NewMethodTemplate(method, name string, params, out map[string]config.Property, isCtx bool) string {

	fn := capitalize(method) + capitalize(name)
	m := fmt.Sprintf(`(r *%sStore)`, capitalize(name))
	p := GetParams(params, isCtx)
	r := GetReturnValues(name, out)

	template := fmt.Sprintf(`func %s %s %s %s`, m, fn, p, r)

	return template
}

func EntityNewFuncTemplate(name string, p, r map[string]config.Property) string {

	var template string
	var funcProcess strings.Builder

	for key := range p {
		funcProcess.WriteString(
			fmt.Sprintf("\t%s:%s,\n", capitalize(key), key),
		)

	}

	fn := NewFuncTemplate("New", name, p, r, true, false)

	contextContent := fmt.Sprintf(`
	ent:=&%s{
		%s
	}
	
	`, GetReturnValues(name, r)[1:], funcProcess.String())

	template = fmt.Sprintf(`
%s {
	%s
	return ent
}
	`, fn, contextContent)

	return template
}

func NewHandlerFuncTemplate() string {

	var fields strings.Builder

	for name := range config.GlobalConfig.Domains {
		fields.WriteString(
			fmt.Sprintf("\t%s: di.%s(db),\n", capitalize(name), capitalize(name)),
		)
	}

	template := fmt.Sprintf(
		`
			func NewHandler(db *db.SQLStore) *Handler {
				return &Handler{
					%s
				}
			}
	
		`, fields.String(),
	)

	return template
}
