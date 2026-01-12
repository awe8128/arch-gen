package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/templates"
	"github.com/awe8128/arch-gen/templates/builder"
	"github.com/awe8128/arch-gen/templates/utils"
)

func GenerateUsecase(domain string, r map[string]config.Repository) (string, string) {
	filename := fmt.Sprintf(`%s.go`, domain)
	content := fmt.Sprintf(`
	%s

	//TODO: Add methods
	%s

	%s

	%s
	`,
		templates.Package(domain),
		templates.Interface(domain, "usecase", domain, nil),
		templates.UsecaseStruct(domain),
		builder.NewFuncBuilder().Name("New", utils.Capitalize(domain), "Usecase").
			AddInProperty(
				"repository",
				fmt.Sprintf("repository.%sRepository", utils.Capitalize(domain)),
				false,
			).
			AddOutProperty(
				domain,
				fmt.Sprintf("%sUsecase", utils.Capitalize(domain)),
				false,
			).Body().BuildFunc(false),
	)
	return content, filename
}
