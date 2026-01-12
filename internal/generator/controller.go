package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates"
	"github.com/awe8128/arch-gen/templates/builder"
	"github.com/awe8128/arch-gen/templates/utils"
)

func GenerateController(domain string, r map[string]config.Repository) (string, string) {
	filename := fmt.Sprintf(`%s.go`, domain)

	content := fmt.Sprintf(`
	%s

	//TODO: Add methods
	%s

	%s

	%s
	`,
		templates.Package("controller"),
		templates.Interface(domain, "controller", domain, nil),
		templates.ControllerStruct(domain),
		builder.NewFuncBuilder().Name("New", utils.Capitalize(domain), "Controller").
			AddInProperty(
				"usecase",
				fmt.Sprintf("%s.%sUsecase", domain, utils.Capitalize(domain)),
				false,
			).
			AddOutProperty(
				"",
				fmt.Sprintf("%sController", utils.Capitalize(domain)),
				false,
			).Body().BuildFunc(false))

	return content, filename
}
