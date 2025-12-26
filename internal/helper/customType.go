package helper

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/internal/templates"
)

func GenerateEntityContent(pkg string, p map[string]config.Property) (string, string, string) {

	content, dir, filename := EntityTemplate(pkg, p)

	return content, dir, filename
}

func GenerateRepositoryContent(pkg string, r map[string]config.Repository) (string, string, string) {

	content, dir, filename := RepositoryTemplate(pkg, r)

	return content, dir, filename
}

func RepositoryTemplate(pkg string, r map[string]config.Repository) (string, string, string) {
	filename := fmt.Sprintf("%s.go", "repository")
	content := fmt.Sprintf(`%s

	%s
	`,
		templates.NewPackageTemplate(pkg),
		templates.NewInterfaceTemplate(pkg, r),
	)

	return content, pkg, filename
}

func EntityTemplate(pkg string, p map[string]config.Property) (string, string, string) {
	filename := fmt.Sprintf("%s.go", pkg)
	content := fmt.Sprintf(`%s

	%s
	%s
	`,
		templates.NewPackageTemplate(pkg),
		templates.NewTypes(pkg, p),
		templates.NewFuncTemplateWithContext("New", pkg, p, nil),
	)

	return content, pkg, filename
}
