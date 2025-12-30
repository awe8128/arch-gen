package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
)

// [name: string, age: uint] -> (name string, age uint)
func GetParams(params map[string]config.Property, addCtx bool) string {
	var p strings.Builder

	for value, custom := range params {

		fieldType := custom.Type
		if custom.Nullable {
			fieldType = "*" + fieldType
		}

		p.WriteString(
			fmt.Sprintf("%s %s,", value, fieldType),
		)
	}

	pStr := p.String()

	// remove last "," and add brackets
	if addCtx {
		return fmt.Sprintf(`(ctx context.Context, %s)`, pStr[:len(pStr)-1])
	} else {
		return fmt.Sprintf(`(%s)`, pStr[:len(pStr)-1])
	}

}

// [name: string, age: uint] -> (name string, age uint)
// nil -> *Name
func GetReturnValues(name string, params map[string]config.Property) string {
	var r strings.Builder

	if params != nil {
		for _, prop := range params {
			fieldType := prop.Type
			if prop.Nullable {
				fieldType = "*" + fieldType
			}

			r.WriteString(
				fmt.Sprintf("%s,", fieldType),
			)
		}

		rStr := r.String()
		return fmt.Sprintf(`(%s)`, rStr[:len(rStr)-1])
	} else {
		return fmt.Sprintf("%s", "*"+capitalize(name))
	}
}
