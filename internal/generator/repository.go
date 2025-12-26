package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/shared/templates"
)

func RepositoryTemplate(pkg string, r map[string]config.Repository) (string, string) {
	filename := fmt.Sprintf("%s.go", "repository")
	content := fmt.Sprintf(`%s

	%s
	`,
		templates.NewPackageTemplate(pkg),
		templates.NewInterfaceTemplate(pkg, r),
	)

	return content, filename
}
