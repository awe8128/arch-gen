package templates

import (
	"fmt"

	"github.com/awe8128/arch-gen/templates/utils"
)

func CreateQueryTemplate(table string, t string, columns string, values string, returns string) string {

	fn := "Create" + utils.Capitalize(table)
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
