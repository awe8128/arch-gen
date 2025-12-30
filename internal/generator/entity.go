package generator

import (
	"fmt"

	"github.com/awe8128/arch-gen/config"
	"github.com/awe8128/arch-gen/shared/templates"
)

func EntityTemplate(name string, p map[string]config.Property) (string, string) {
	filename := fmt.Sprintf("%s.go", name)

	content := fmt.Sprintf(`%s

	%s
	%s
	`,
		templates.NewPackageTemplate(name),
		templates.NewStructTemplate(name, p),
		templates.EntityNewFuncTemplate(name, p, nil),
	)

	return content, filename
}
