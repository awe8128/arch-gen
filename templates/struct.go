package templates

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates/builder"
	"github.com/awe8128/arch-gen/templates/utils"
)

func NewStructTemplate(prefix, name, suffix string, p map[string]config.Property) string {
	b := builder.NewStructBuilder().Name(utils.Capitalize(name), prefix, suffix)
	for fieldName, prop := range p {
		b.AddProperties(utils.Capitalize(fieldName), prop.Type, prop.Nullable)
	}

	template := b.Build()

	return template
}

func UsecaseStructTemplate(name string) string {
	b := builder.NewStructBuilder().Name(name, "", "Usecase")

	b.AddProperties(
		"repository",
		fmt.Sprintf("repository.%sRepository", utils.Capitalize(name)),
		false,
	)

	template := b.Build()

	return template
}

func ControllerStructTemplate(name string) string {
	b := builder.NewStructBuilder().Name(name, "", "Controller")

	b.AddProperties(
		"usecase",
		fmt.Sprintf("%s.%sUsecase", name, utils.Capitalize(name)),
		false,
	)

	template := b.Build()

	return template
}

func HandlerStructTemplate() string {
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
