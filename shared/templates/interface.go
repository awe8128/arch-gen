package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates/immutable"
)

func InterfaceMethodTemplate(name string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := immutable.NewFuncBuilder(
			method, capitalize(name),
			fn.In,
			fn.Out,
		).Body().Method(capitalize(name) + "Store").BuildFunc(true)

		fields.WriteString(
			fmt.Sprintf("\t%s\n", fnStr),
		)
	}

	return fields.String()

}

/*
Interface Template Creates

	type <name> interface {
		func()
	}
*/
func InterfaceTemplate(name, suffix, domain string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := immutable.NewFuncBuilder(method, capitalize(domain), fn.In, fn.Out).BuildInterface(true)
		fields.WriteString(
			fmt.Sprintf("\t%s\n", fnStr),
		)
	}

	interfaceName := capitalize(name) + capitalize(suffix)

	content := fmt.Sprintf(
		`
	type %s interface {
	%s}
	`,
		interfaceName, fields.String(),
	)

	return content
}
