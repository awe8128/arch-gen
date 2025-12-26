package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/shared/templates"
)

func EntityTemplate(pkg string, p map[string]config.Property) (string, string) {
	filename := fmt.Sprintf("%s.go", pkg)
	content := fmt.Sprintf(`%s

	%s
	%s
	`,
		templates.NewPackageTemplate(pkg),
		templates.NewStructTemplate(pkg, p),
		templates.NewFuncTemplateWithContext("New", pkg, p, nil),
	)

	return content, filename
}
