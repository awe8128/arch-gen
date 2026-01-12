package generator

import (
	"strconv"

	"github.com/awe8128/arch-gen/templates"
)

func GenerateMigration(table string, columns map[string]string, tag int) (string, string) {

	migrateTag := "00000"

	prefix := generateTag(tag, migrateTag)

	filename := generateFileName(prefix, table, "up")

	content := templates.Migration(table, columns)

	return content, filename

}

func generateTag(tag int, base string) string {
	tagStr := strconv.Itoa(tag)
	return base + tagStr
}

func generateFileName(prefix, name, method string) string {

	res := prefix + "_" + name + "." + method + ".sql"

	return res
}
