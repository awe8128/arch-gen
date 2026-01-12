package generator

import "github.com/awe8128/arch-gen/templates"

func SqlcYamlTemplate() (string, string) {
	filename := "sqlc.yaml"
	content := templates.NewSQLCYamlTemplate()
	return content, filename
}
