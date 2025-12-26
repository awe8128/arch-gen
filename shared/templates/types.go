package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

func NewTypes(name string, p map[string]config.Property) string {
	var fields strings.Builder

	for fieldName, prop := range p {
		fieldType := prop.Type
		if prop.Nullable {
			fieldType = "*" + fieldType
		}

		fields.WriteString(
			fmt.Sprintf("\t%s %s\n", capitalize(fieldName), fieldType),
		)
	}

	content := fmt.Sprintf(
		`
type %s struct {
%s}
`,
		capitalize(name),
		fields.String(),
	)

	return content
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
