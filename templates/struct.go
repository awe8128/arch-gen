package templates

import (
	"fmt"
	"strings"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates/builder"
	"github.com/awe8128/arch-gen/templates/utils"
)

func NewStruct(prefix, name, suffix string, p map[string]config.Property) string {
	b := builder.NewStructBuilder().Name(utils.Capitalize(name), prefix, suffix)
	for fieldName, prop := range p {
		b.AddProperties(utils.Capitalize(fieldName), prop.Type, prop.Nullable)
	}

	template := b.Build()

	return template
}

func UsecaseStruct(name string) string {
	b := builder.NewStructBuilder().Name(name, "", "Usecase")

	b.AddProperties(
		"repository",
		fmt.Sprintf("repository.%sRepository", utils.Capitalize(name)),
		false,
	)

	template := b.Build()

	return template
}

func ControllerStruct(name string) string {
	b := builder.NewStructBuilder().Name(name, "", "Controller")

	b.AddProperties(
		"usecase",
		fmt.Sprintf("%s.%sUsecase", name, utils.Capitalize(name)),
		false,
	)

	template := b.Build()

	return template
}

func HandlerStruct() string {
	b := builder.NewStructBuilder().Name("Handler", "", "")
	for name := range config.GlobalConfig.Domains {
		b.AddProperties(
			utils.Capitalize(name),
			fmt.Sprintf("controller.%sController", utils.Capitalize(name)),
			false,
		)

	}
	template := b.Build()

	return template
}

var configs = []string{
	"DB_SOURCE",
	"DB_NAME",
	"DB_USER",
	"DB_PASSWORD",
	"DB_PORT",
	"DB_HOST",
	"API_PORT",
}

func ConfigStruct() string {
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
