package templates

import "fmt"

func CreateQueryTemplate(table string, t string, columns string, values string, returns string) string {

	fn := "Create" + capitalize(table)
	dbTable := table + "s"

	content := fmt.Sprintf(`
-- name: %s :%s
INSERT INTO %s (
	%s
) VALUES (
	%s
)
RETURNING %s;
	`, fn, t, dbTable, columns, values, returns,
	)

	return content
}
