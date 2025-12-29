package generator

import (
	"strconv"

	"github.com/awe8128/arch-gen/internal/templates"
)

func MigrationTemplate(table string, columns map[string]string, tag int) (string, string) {

	migrateTag := "00000"

	prefix := generateTag(tag, migrateTag)

	filename := generateMigrateFile(prefix, table, "up")

	content := templates.MigrationTemplate(table, columns)

	return content, filename

}

func generateTag(tag int, base string) string {
	tagStr := strconv.Itoa(tag)
	return base + tagStr
}

func generateMigrateFile(prefix, name, method string) string {

	res := prefix + "_" + name + "." + method + ".sql"

	return res
}
