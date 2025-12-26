package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/shared/templates"
)

func InfraRepositoryTemplate(pkg string, r map[string]config.Repository, p map[string]config.Property) (string, string) {
	filename := fmt.Sprintf("%s.go", pkg)
	content := fmt.Sprintf(`%s

	%s
	%s
	`,
		templates.NewPackageTemplate("repository"),

		templates.InterfaceTemplateWithName(pkg, r),

		templates.NewInfraTemplate(pkg),
	)

	return content, filename
}
