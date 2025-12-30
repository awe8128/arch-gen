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
	%s
	`,
		templates.NewPackageTemplate("repository"),

		templates.InterfaceTemplate(pkg, r),

		templates.NewStoreTemplate(pkg),

		templates.NewInterfaceMethod(pkg, r),
	)

	return content, filename
}
