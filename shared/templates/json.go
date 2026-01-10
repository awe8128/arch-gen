package templates

import (
	"fmt"
	"strings"
)

var configs = []string{
	"DB_SOURCE",
	"DB_NAME",
	"DB_USER",
	"DB_PASSWORD",
	"DB_PORT",
	"DB_HOST",
	"API_PORT",
}

func NewJSONStruct() string {
	var fields strings.Builder
	for _, config := range configs {
		tag := fmt.Sprintf(`mapstructure:"%s"`, config)

		jsonTag := fmt.Sprintf("`%s`", tag)
		fields.WriteString(
			fmt.Sprintf("\t%s %s %s\n", config, "string", jsonTag),
		)
	}

	template := fmt.Sprintf(`
	type Config struct {
		%s
	}
	`, fields.String())

	return template
}
