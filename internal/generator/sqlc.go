package generator

import "github.com/awe8128/arch-gen/shared/templates"

func SqlcYamlTemplate() (string, string) {
	filename := "sqlc.yaml"
	content := templates.NewSqlcYamlTemplate()
	return content, filename
}
