package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates/builder"
	"github.com/awe8128/arch-gen/templates/utils"
)

func InterfaceMethod(name string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := builder.NewFuncBuilder().
			Name(utils.Capitalize(method), utils.Capitalize(name), "").
			In(fn.In).Out(fn.Out).
			Body().Method(utils.Capitalize(name) + "Store").BuildFunc(true)

		fields.WriteString(
			fmt.Sprintf("\t%s\n", fnStr),
		)
	}

	return fields.String()

}

func Interface(name, suffix, domain string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := builder.NewFuncBuilder().
			Name(utils.Capitalize(method), utils.Capitalize(domain), "").
			In(fn.In).Out(fn.Out).
			BuildInterface(true)
		fields.WriteString(
			fmt.Sprintf("\t%s\n", fnStr),
		)
	}

	interfaceName := utils.Capitalize(name) + utils.Capitalize(suffix)

	content := fmt.Sprintf(
		`
	type %s interface {
	%s}
	`,
		interfaceName, fields.String(),
	)

	return content
}
