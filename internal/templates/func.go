package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

// Create(ctx context.Context, ...) (types, types)
func NewFuncTemplate(method, name string, in, out map[string]config.Property) string {
	var content string
	var outContent string
	var inFields strings.Builder
	var outFields strings.Builder

	for key, prop := range in {
		fieldType := prop.Type
		if prop.Nullable {
			fieldType = "*" + fieldType
		}

		inFields.WriteString(
			fmt.Sprintf("%s %s,", key, fieldType),
		)
	}

	inParams := inFields.String()
	inContent := fmt.Sprintf(`%s(ctx context.Context, %s)`, capitalize(method)+capitalize(name), inParams[:len(inParams)-1])

	if out != nil {
		for _, prop := range out {
			fieldType := prop.Type
			if prop.Nullable {
				fieldType = "*" + fieldType
			}

			outFields.WriteString(
				fmt.Sprintf("%s,", fieldType),
			)
		}

		outParams := outFields.String()
		outContent = fmt.Sprintf(`(%s)`, outParams[:len(outParams)-1])
	} else {
		outContent = fmt.Sprintf("%s", "*"+capitalize(name))
	}

	content = inContent + outContent

	return content
}

//	func Create(ctx context.Context, ...) (types, types) || *entity {
//		 return nil
//	}
func NewFuncTemplateWithContext(method, name string, in, out map[string]config.Property) string {
	var template string
	var outContent string

	var inFields strings.Builder
	var outFields strings.Builder
	var contextField strings.Builder

	for key, prop := range in {
		fieldType := prop.Type
		if prop.Nullable {
			fieldType = "*" + fieldType
		}

		inFields.WriteString(
			fmt.Sprintf("%s %s,", key, fieldType),
		)

		contextField.WriteString(
			fmt.Sprintf("\t%s:%s,\n", capitalize(key), key),
		)

	}

	inParams := inFields.String()
	inContent := fmt.Sprintf(`%s(ctx context.Context, %s)`, capitalize(method)+capitalize(name), inParams[:len(inParams)-1])

	if out != nil {
		for _, prop := range out {
			fieldType := prop.Type
			if prop.Nullable {
				fieldType = "*" + fieldType
			}

			outFields.WriteString(
				fmt.Sprintf("%s,", fieldType),
			)
		}

		outParams := outFields.String()
		outContent = fmt.Sprintf(`(%s)`, outParams[:len(outParams)-1])
	} else {
		outContent = fmt.Sprintf(`%s`, "*"+capitalize(name))
	}

	contextContent := fmt.Sprintf(`
	ent:=&%s{
		%s
	}
	
	`, outContent[1:], contextField.String())

	// TODO: add more context for typical functions
	template = fmt.Sprintf(`
	func %s {
		%s
		return ent
	}
	`, inContent+outContent, contextContent)

	return template
}
