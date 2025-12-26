package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

func NewInterfaceTemplate(name string, r map[string]config.Repository) string {
	var fields strings.Builder

	for method, fn := range r {
		fnStr := NewFuncTemplate(method, name, fn.In, fn.Out)

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
