package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

func NewInterfaceMethod(name string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := NewMethodTemplate(method, name, fn.In, fn.Out, true)

		fields.WriteString(
			fmt.Sprintf("\t%s\n", fnStr),
		)
	}

	return fields.String()

}

func NewInterfaceTemplate(name string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := NewFuncTemplate(method, name, fn.In, fn.Out, true, true)

		fields.WriteString(
			fmt.Sprintf("\t%s\n", fnStr),
		)
	}

	content := fmt.Sprintf(
		`
type Repository interface {
%s}
`,
		fields.String(),
	)

	return content
}

/*
InterfaceTemplate creates
type <name> interface {

}
*/
func InterfaceTemplate(name, suffix string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := NewFuncTemplate(method, name, fn.In, fn.Out, true, true)

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
