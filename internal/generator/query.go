package generator

import (
	"fmt"

	"strconv"
	"strings"

	"github.com/awe8128/arch-gen/templates"
)

func GenerateQuery(table string, columns map[string]string) (string, string) {
	var fields strings.Builder
	var values strings.Builder

	var tag = 1
	for column := range columns {
		fields.WriteString(
			fmt.Sprintf("\t%s,\n", column),
		)

		val := "$" + strconv.Itoa(tag) + ","
		tag++
		values.WriteString(val)

	}
	f := fields.String()
	v := values.String()

	filename := table + ".sql"
	content := templates.CreateQuery(table, "one", f[:len(f)-2], v[:len(v)-1], "*")

	return content, filename
}
