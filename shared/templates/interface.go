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

func InterfaceTemplate(name string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := NewFuncTemplate(method, name, fn.In, fn.Out, true, true)

		fields.WriteString(
			fmt.Sprintf("\t%s\n", fnStr),
		)
	}

	content := fmt.Sprintf(
		`
type %s interface {
%s}
`,
		fmt.Sprintf("%sRepository", capitalize(name)), fields.String(),
	)

	return content
}
