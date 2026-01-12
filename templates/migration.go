package templates

import (
	"fmt"
	"strings"
)

func Migration(name string, columns map[string]string) string {

	var fields strings.Builder

	for column, t := range columns {
		col := `"` + column + `"`
		fields.WriteString(
			fmt.Sprintf("\t%s %s,\n", col, t),
		)
	}

	f := fields.String()

	template := fmt.Sprintf(
		`
CREATE TABLE "%s" (

%s
);
`,
		name+"s", f[:len(f)-2],
	)

	return template
}
